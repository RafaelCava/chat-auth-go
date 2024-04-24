package controllers

import (
	"net/http"
	"sync"

	"github.com/RafaelCava/chat-auth-go/domain/usecases"
	"github.com/RafaelCava/chat-auth-go/presentation/protocols"
	"github.com/gin-gonic/gin"
)

type CreateRoomController struct {
	CreateRoomUseCase usecases.CreateRoomUseCase
}

func NewCreateRoomController(createRoomUseCase usecases.CreateRoomUseCase) protocols.Controller {
	return &CreateRoomController{createRoomUseCase}
}

func (controller *CreateRoomController) Handle(ctx *gin.Context) error {
	var params usecases.CreateRoomParams
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
		result, err := controller.CreateRoomUseCase.Execute(params)
		if err != nil {
			resultChannel <- protocols.HttpResponse{StatusCode: http.StatusBadRequest, Body: err.Error()}
		} else {
			resultChannel <- protocols.HttpResponse{StatusCode: http.StatusCreated, Body: result}
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
