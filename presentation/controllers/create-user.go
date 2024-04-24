package controllers

import (
	"net/http"
	"sync"

	"github.com/RafaelCava/chat-auth-go/domain/usecases"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	createUserUseCase usecases.CreateUserUseCase
}

func NewCreateUserController(createUserUseCase usecases.CreateUserUseCase) protocols.Controller {
	return &CreateUserController{createUserUseCase}
}

func (controller *CreateUserController) Handle(ctx *gin.Context) error {
	var params usecases.CreateUserParams
	err := ctx.BindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	if err := params.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	resultChannel := make(chan protocols.HttpResponse, 1)
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		user, err := controller.createUserUseCase.Execute(params)
		if err != nil {
			resultChannel <- protocols.HttpResponse{StatusCode: http.StatusBadRequest, Body: err.Error()}
		} else {
			resultChannel <- protocols.HttpResponse{StatusCode: http.StatusCreated, Body: user}
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
