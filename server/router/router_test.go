package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/notes/config"
	"github.com/stretchr/testify/assert"
)

func TestNewGinRouter(t *testing.T) {

	cfg := &config.Configuration{
		EnvironmentName: "test",
		Server: config.Server{
			RestServicePort: 8080,
			ReadTimeout:     10,
			WriteTimeout:    10,
			IdleTimeout:     10,
		},
		Swagger:  config.Swagger{},
		Postgres: config.DB{},
	}
	g, err := NewGinRouter(cfg)
	assert.NoError(t, err, "NewGinRouter should not return an error")

	assert.NotNil(t, g, "NewGinRouter should return a non-nil *gin.Engine instance")
	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})
	w := performRequest(g, "GET", "/test")
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "test", "Expected response body to contain 'test'")
}

func performRequest(r *gin.Engine, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
