package protocols

import "github.com/RafaelCava/chat-auth-go/domain/models"

type CreateUserRepository interface {
	CreateUser(username, password, email, logo string) (*models.User, error)
}
