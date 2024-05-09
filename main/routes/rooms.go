package routes

import (
	"os"
	"strings"

	"github.com/RafaelCava/chat-auth-go/infra/cryptography/adapters"
	"github.com/RafaelCava/chat-auth-go/main/factories"
	"github.com/gin-gonic/gin"
)

func addRoomsRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/rooms")
	users.POST("", middlewareValidation, func(ctx *gin.Context) {
		factories.NewCreateRoomControllerFactory().Handle(ctx)
	})
}

func middlewareValidation(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "No token provided"})
		return
	}
	splitToken := strings.Split(authorizationHeader, " ")
	prefix := splitToken[0]
	if prefix != "Bearer" {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token type"})
		return
	}
	token := splitToken[1]
	if token == "" || token == "null" {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token credential"})
		return
	}
	secretKey := os.Getenv("SECRET_TOKEN")
	issuerKey := os.Getenv("ISS_TOKEN")
	encrypterAdapter := adapters.NewEncrypterAdapter(secretKey, issuerKey)
	claims, err := encrypterAdapter.Decrypt(token)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
	}
	ctx.Set("user_id", claims.UserId)
	ctx.Next()
}
