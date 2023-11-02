package tests

import (
	"albert/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthController(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()

	router.GET("/ping", controllers.Ping)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
