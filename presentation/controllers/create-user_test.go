package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/RafaelCava/chat-auth-go/domain/usecases"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type CreateUserUseCaseSpy struct {
	options map[string]interface{}
}

func (spy *CreateUserUseCaseSpy) Execute(params usecases.CreateUserParams) (*models.User, error) {
	if spy.options["error"] == true {
		return nil, errors.New("error")
	}
	if spy.options["returnNil"] == true {
		return nil, nil
	}
	return &models.User{
		ID:        "any_id",
		Username:  "any_username",
		Logo:      "any_logo",
		Email:     "any_email",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
func TestCreateUserController(t *testing.T) {
	t.Run("should return 400 if invalid body is provided", func(t *testing.T) {
		sut := controllers.NewCreateUserController(&CreateUserUseCaseSpy{
			options: map[string]interface{}{},
		})
		router := gin.Default()
		router.POST("/users", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader("invalid json")
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", reqBody)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 if Validate returns an error", func(t *testing.T) {
		sut := controllers.NewCreateUserController(&CreateUserUseCaseSpy{
			options: map[string]interface{}{},
		})

		router := gin.Default()
		router.POST("/users", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader(`{"username": "any_username", "email": "any_email"}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", reqBody)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 if CreateUserUseCase returns an error", func(t *testing.T) {
		sut := controllers.NewCreateUserController(&CreateUserUseCaseSpy{
			options: map[string]interface{}{"error": true},
		})

		router := gin.Default()
		router.POST("/users", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader(`{"username": "any_username", "email": "teste@teste.com", "password": "any_password"}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", reqBody)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 201 if CreateUserUseCase returns a user", func(t *testing.T) {
		sut := controllers.NewCreateUserController(&CreateUserUseCaseSpy{
			options: map[string]interface{}{},
		})

		router := gin.Default()
		router.POST("/users", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader(`{"username": "any_username", "email": "teste@teste.com", "password": "any_password"}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", reqBody)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
