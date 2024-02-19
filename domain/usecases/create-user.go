package usecases

import "github.com/RafaelCava/chat-auth-go/domain/models"

type CreateUserParams struct {
	Username string
	Password string
	Email    string
	Logo     string
}
type CreateUserUseCase interface {
	Execute(params CreateUserParams) (*models.User, error)
}
