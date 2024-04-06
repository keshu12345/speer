package notesService

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mocksauth "github.com/keshu12345/notes/mocks/auth"
	mocknotes "github.com/keshu12345/notes/mocks/notes"
	"github.com/keshu12345/notes/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Assume these mock services are already defined as in the previous example
// type MockAuthService struct{}
// type MockNotesService struct{}

func TestRegisterNotesformEndPoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockAuthService := mocksauth.NewAuthService(t)
	mockNotesService := mocknotes.NewNotesService(t)
	user := model.User{}
	// notes := []model.Note{}
	// note := model.Note{}

	mockAuthService.On("SingUpUser", mock.AnythingOfType("*gin.Context")).Return(user, nil)
	mockAuthService.On("SingInUser", mock.AnythingOfType("*gin.Context")).Return(mock.Anything, nil)
	// mockNotesService.On("GetAllNotes", mock.AnythingOfType("*gin.Context")).Return(notes, nil)
	// mockNotesService.On("GetNoteByID", mock.AnythingOfType("*gin.Context")).Return(note, nil)
	RegisterNotesformEndPoint(router, mockAuthService, mockNotesService)

	t.Run("SignUp Route", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/signup", nil)
		router.ServeHTTP(w, req)

		// Adjust the expected status code based on your signUp handler's implementation
		assert.Equal(t, http.StatusCreated, w.Code, "Expected /api/signup to return 200")
	})

	t.Run("SignIn Route", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/signin", nil)
		router.ServeHTTP(w, req)

		// Adjust the expected status code based on your signIn handler's implementation
		assert.Equal(t, http.StatusOK, w.Code, "Expected /api/signin to return 200")
	})

	// t.Run("GetAllNotes Route - Unauthorized", func(t *testing.T) {
	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest("GET", "/api/notes/", nil)
	// 	req.Header.Set("Authorization", "Bearer mock-valid-token")
	// 	router.ServeHTTP(w, req)

	// 	// Assuming authMiddleware denies request without valid token
	// 	assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected /api/notes/ to return 401 for unauthorized access")
	// })

	// t.Run("GetAllNotes Route - Authorized", func(t *testing.T) {
	// 	// Mock a request with an "Authorization" header, assuming a Bearer token scheme
	// 	// You will need to adjust this based on your actual authMiddleware implementation
	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest("GET", "/api/notes/", nil)
	// 	req.Header.Set("Authorization", "Bearer mock-valid-token")
	// 	router.ServeHTTP(w, req)

	// 	// Adjust the expected status code based on your GetAllNotesRecords handler's implementation and authMiddleware's behavior
	// 	assert.Equal(t, http.StatusOK, w.Code, "Expected /api/notes/ to return 200 for authorized access")
	// })

	// t.Run("GetNoteById Route - Authorized", func(t *testing.T) {
	// 	// Mock a request to get a note by ID, assuming a note exists with ID 1
	// 	// Again, this test assumes the presence of an "Authorization" header for simplicity
	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest("GET", "/api/notes/1", nil)
	// 	req.Header.Set("Authorization", "Bearer mock-valid-token")
	// 	router.ServeHTTP(w, req)

	// 	// Adjust the expected status code based on your GetNotesByIdRecords handler's implementation and authMiddleware's behavior
	// 	assert.Equal(t, http.StatusOK, w.Code, "Expected /api/notes/1 to return 200 for authorized access")
	// })

	// Add more tests as needed for other routes and scenarios
}
