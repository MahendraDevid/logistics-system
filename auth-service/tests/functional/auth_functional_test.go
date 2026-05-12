package functional

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth-service/internal/handler"
	"auth-service/mocks"
	"auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthFlow(t *testing.T) {
	// 1. Setup Mock & Service
	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)
	authHandler := handler.NewAuthHandler(authService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/auth/register", authHandler.Register)

	// 2. Mock Behavior (Auto-respond)
	mockRepo.On("Create", mock.Anything).Return(nil)

	// 3. Test Request Register
	userData := map[string]string{
		"name":     "Imam UPI",
		"email":    "imam@upi.edu",
		"password": "password123",
		"role":     "kurir",
	}
	body, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	
	r.ServeHTTP(w, req)

	// 4. Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Registrasi berhasil")
}