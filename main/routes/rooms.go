package routes

import (
	"github.com/RafaelCava/chat-auth-go/main/factories"
	"github.com/gin-gonic/gin"
)

func addRoomsRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/rooms")
	users.POST("", func(ctx *gin.Context) {
		factories.NewCreateRoomControllerFactory().Handle(ctx)
	})
}
