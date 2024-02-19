package protocols

import "github.com/RafaelCava/chat-auth-go/domain/models"

type FindUserByEmailRepository interface {
	FindByEmail(email string, projection []string) (*models.User, error)
}
