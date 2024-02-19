package routes

import (
	"github.com/RafaelCava/chat-auth-go/main/factories"
	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	auth.POST("/login", func(ctx *gin.Context) {
		factories.NewLoginControllerFactory().Handle(ctx)
	})
}
