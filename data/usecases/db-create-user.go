package usecases

import (
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/RafaelCava/chat-auth-go/domain/usecases"
)

type DbCreateUserUseCase struct{}

func NewDbCreateUserUseCase() usecases.CreateUserUseCase {
	return &DbCreateUserUseCase{}
}

func (useCase *DbCreateUserUseCase) Execute(params usecases.CreateUserParams) (*models.User, error) {
	user, err := models.NewUser(params.Username, params.Password, params.Email, params.Logo)
	if err != nil {
		return nil, err
	}
	return user, nil
}
