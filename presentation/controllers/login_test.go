package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/RafaelCava/chat-auth-go/domain/usecases"
	"github.com/RafaelCava/chat-auth-go/presentation/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type AuthUseCaseSpy struct {
	options OptionsSpy
}

func (spy *AuthUseCaseSpy) Execute(params usecases.AuthRequest) (*models.TokenPair, error) {
	spy.options.count = spy.options.count + 1
	if spy.options.returnError == true {
		return nil, errors.New("error")
	}
	if spy.options.returnNil == true {
		return nil, nil
	}
	return &models.TokenPair{
		AccessToken:  "any_access_token",
		RefreshToken: "any_refresh_token",
	}, nil
}

func TestLoginController(t *testing.T) {
	t.Run("should return 400 if invalid body is provided", func(t *testing.T) {
		spy := AuthUseCaseSpy{}
		sut := controllers.NewLoginController(&spy)

		router := gin.Default()
		router.POST("/auth/login", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader("invalid json")
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", reqBody)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, 0, spy.options.count)
	})

	t.Run("should return 409 if invalid body is provided", func(t *testing.T) {
		spy := AuthUseCaseSpy{}
		sut := controllers.NewLoginController(&spy)

		router := gin.Default()
		router.POST("/auth/login", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader(`{"username": "any_username", "password": "any_password"}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", reqBody)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Equal(t, 0, spy.options.count)
	})

	t.Run("should return 401 if invalid credentials are provided", func(t *testing.T) {
		spy := AuthUseCaseSpy{OptionsSpy{returnError: true}}
		sut := controllers.NewLoginController(&spy)

		router := gin.Default()
		router.POST("/auth/login", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader(`{"email": "teste@teste.com", "password": "any_password"}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", reqBody)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, 1, spy.options.count)
	})

	t.Run("should return 200 if valid credentials are provided", func(t *testing.T) {
		spy := AuthUseCaseSpy{}
		sut := controllers.NewLoginController(&spy)

		router := gin.Default()
		router.POST("/auth/login", func(ctx *gin.Context) {
			sut.Handle(ctx)
		})

		reqBody := strings.NewReader(`{"email": "teste@teste.com", "password": "any_password"}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", reqBody)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 1, spy.options.count)
		assert.Equal(t, "{\"data\":{\"access_token\":\"any_access_token\",\"refresh_token\":\"any_refresh_token\"}}", w.Body.String())
	})
}
