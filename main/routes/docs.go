package routes

import (
	_ "github.com/RafaelCava/chat-auth-go/main/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func addDocsRoutes(rg *gin.RouterGroup) {
	rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
