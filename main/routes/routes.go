package routes

import (
	"context"
	"net/http"
	"os"

	"github.com/RafaelCava/chat-auth-go/main/factories"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

var router = gin.New()
var server = &http.Server{
	Addr:    ":3000",
	Handler: router,
}

func Run() error {
	getRoutes()
	port := os.Getenv("PORT")
	if port != "" {
		server.Addr = ":" + port
	}
	return server.ListenAndServe()
}

func ShutDown(ctx context.Context) error {
	return server.Shutdown(ctx)
}

func getRoutes() {
	apiPrefix := router.Group("/api")
	apiPrefix.Use(apmgin.Middleware(router))
	apiPrefix.Use(factories.NewCorsMiddleware())
	apiPrefix.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	addUserRoutes(apiPrefix)
	// addDocsRoutes(apiPrefix)
	// addHealthCheckRoutes(apiPrefix)
	// addAuthRoutes(apiPrefix)
}
