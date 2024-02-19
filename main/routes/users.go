package routes

import (
	"github.com/RafaelCava/chat-auth-go/main/factories"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.POST("", func(ctx *gin.Context) {
		factories.NewCreateUserControllerFactory().Handle(ctx)
	})
}
