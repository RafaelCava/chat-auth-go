package usecases

import "github.com/RafaelCava/chat-auth-go/domain/models"

type CreateUserUseCase interface {
	Execute(username, password, email, logo string) (*models.User, error)
}
