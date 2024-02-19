package controllers

import (
	"net/http"
	"sync"

	"github.com/RafaelCava/chat-auth-go/domain/usecases"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	authUseCase usecases.AuthUseCase
}

func NewLoginController(authUseCase usecases.AuthUseCase) protocols.Controller {
	return &LoginController{authUseCase}
}

func (controller *LoginController) Handle(ctx *gin.Context) error {
	var request usecases.AuthRequest
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
	}
	resultChannel := make(chan protocols.HttpResponse, 1)
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		tokenPair, err := controller.authUseCase.Execute(request)
		if err != nil {
			resultChannel <- protocols.HttpResponse{StatusCode: http.StatusBadRequest, Body: err.Error()}
		} else {
			resultChannel <- protocols.HttpResponse{StatusCode: http.StatusOK, Body: tokenPair}
		}
	}()
	wg.Add(1)
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	result := <-resultChannel
	ctx.JSON(result.StatusCode, gin.H{"data": result.Body})
	return nil
}
