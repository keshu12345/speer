package notesService

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/keshu12345/notes/model"
	logger "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// mockery --exported --dir=./ -name=AuthService --case underscore -output ../mocks/auth
type AuthService interface {
	SingUpUser(c *gin.Context) (model.User, error)
	SingInUser(c *gin.Context) (string, error)
}

type NotesService interface {
	GetAllNotes() ([]model.Note, error)
	GetNoteByID(noteId uint64) (model.Note, error)
}

type getNotesService struct {
	fx.In
	Client *gorm.DB
}

type getAuthService struct {
	// gorm dependency injected
	fx.In
	// gorm client
	Client *gorm.DB
}

func NewGetNotesService(getNoteServiceStruct getNotesService) NotesService {
	return getNotesService{Client: getNoteServiceStruct.Client}
}

func NewGetAuthService(gasStruct getAuthService) AuthService {
	return getAuthService{Client: gasStruct.Client}
}

func (getAuthService getAuthService) SingUpUser(c *gin.Context) (model.User, error) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Error("Unable to bind user %w", err)
		return user, fmt.Errorf("Unable to bind user %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password %w", err)
		return user, fmt.Errorf("Failed to hash password %w", err)
	}

	user.Password = string(hashedPassword)

	if err := getAuthService.Client.Create(&user).Error; err != nil {
		logger.Error("Failed to hash password %w", err)
		return user, fmt.Errorf("Failed to create user %w", err)
	}

	return user, nil
}
func (authService getAuthService) SingInUser(c *gin.Context) (string, error) {

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", fmt.Errorf("Unable to bind user %v", err)
	}

	var storedUser model.User
	if err := authService.Client.Where("email = ?", user.Email).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return "", fmt.Errorf("error:Invalid credentials %w", err)
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return "", fmt.Errorf("error:Invalid credentials %w", err)
	}

	token, err := createToken(storedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return "", fmt.Errorf("error:Failed to generate token %w", err)
	}

	return token, nil

}

func createToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (noteService getNotesService) GetAllNotes() ([]model.Note, error) {
	var notes []model.Note

	// Try fetching notes without Preload
	if err := noteService.Client.Find(&notes).Error; err != nil {
		logger.Fatalf("Failed to fetch notes directly: %v", err)
		return nil, err
	}

	// If direct fetch works, proceed to Preload
	if err := noteService.Client.Preload("User").Find(&notes).Error; err != nil {
		logger.Fatalf("Failed to fetch notes with Preload: %v", err)
		return nil, err
	}

	return notes, nil
}

func (noteService getNotesService) GetNoteByID(noteID uint64) (model.Note, error) {
	var note model.Note

	if err := noteService.Client.Preload("User").First(&note, noteID).Error; err != nil {
		logger.Error("Failed to find note by ID: %w", err)
		return model.Note{}, err
	}

	return note, nil
}
