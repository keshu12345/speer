package notesService

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// var db *gorm.DB

var jwtKey = []byte("speer")

// type Employee struct {
// 	gorm.Model
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"-"`
// 	// Add other relevant fields as needed
// }

var authServiceInterface AuthService
var noteServiceInterface NotesService

// Endpoint
func RegisterNotesformEndPoint(g *gin.Engine, authService AuthService, noteService NotesService) {
	authServiceInterface = authService
	g.POST("/api/signup", signUp)
	g.POST("/api/signin", signIn)
	noteRecord := g.Group("/api/notes")
	{
		noteServiceInterface = noteService
		noteRecord.Use(authMiddleware)
		noteRecord.GET("/", GetAllNotesRecords)
		noteRecord.GET("/:id", GetNotesByIdRecords)
		authServiceInterface = authService

		// gcpsInterafce = gcps
		// healthCheckInterface = hcs

	}
	// g.GET("/healthcheck", registerHealthCheck)
	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// func GetAllNotesRecords(c *gin.Context) {
// 	// authServiceInterface.SingUpUser()

// }

func GetNotesByIdRecords(c *gin.Context) {
	// noteID := extractNoteIDFromContext(c) // Implement function to extract note ID from context

	noteIDStr := c.Param("noteID")

	// Convert the noteIDStr to uint (assuming note ID is uint)
	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		// Handle error for invalid noteID format
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	// Use noteID (uint) for further processing
	// Call the GetNoteByID function passing noteID
	note, err := noteServiceInterface.GetNoteByID(noteID)
	if err != nil {
		// Handle error response for GetNoteByID
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch note"})
		return
	}

	// Return the note in the response
	c.JSON(http.StatusOK, gin.H{"note": note})

}
func GetAllNotesRecords(c *gin.Context) {
	notes, err := noteServiceInterface.GetAllNotes()
	if err != nil {
		// Handle error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}
	// Return notes in the response
	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func signUp(c *gin.Context) {

	user, err := authServiceInterface.SingUpUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"user": user})
	}

}
func signIn(c *gin.Context) {
	token, err := authServiceInterface.SingInUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized access",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
