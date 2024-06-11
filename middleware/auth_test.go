package middleware_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example/hello/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name             string
		username         string
		password         string
		expectedCode     int
		expectedResponse gin.H
	}{
		{
			name:             "No credentials",
			username:         "",
			password:         "",
			expectedCode:     http.StatusUnauthorized,
			expectedResponse: gin.H{"error": "Authorization basic token required"},
		},
		{
			name:             "Invalid credentials",
			username:         "wronguser",
			password:         "wrongpass",
			expectedCode:     http.StatusUnauthorized,
			expectedResponse: gin.H{"error": "Invalid authorization token"},
		},
		{
			name:             "Valid credentials",
			username:         "useradmin123",
			password:         "useradmin123",
			expectedCode:     http.StatusOK,
			expectedResponse: gin.H{"status": "success"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create gin router
			r := gin.New()
			r.Use(middleware.AuthMiddleware())

			// Test middleware
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "success"})
			})

			// Create new HTTP request
			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tt.username != "" && tt.password != "" {
				req.SetBasicAuth(tt.username, tt.password)
			}

			// Record the response
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Check the response
			assert.Equal(t, tt.expectedCode, w.Code)
			var response gin.H
			// err := gin.BindJSON(w.Body.Bytes(), &response)
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResponse, response)
		})
	}
}

