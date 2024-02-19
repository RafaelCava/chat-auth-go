package usecases

import (
	"errors"

	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/RafaelCava/chat-auth-go/domain/usecases"
)

type DbCreateUserUseCase struct {
	validateIfUserExistsRepository protocols.ValidateIfUserExistsRepository
	hasher                         protocols.Hasher
	createUserRepository           protocols.CreateUserRepository
}

func NewDbCreateUserUseCase(validateIfUserExistsRepository protocols.ValidateIfUserExistsRepository, hasher protocols.Hasher, createUserRepository protocols.CreateUserRepository) usecases.CreateUserUseCase {
	return &DbCreateUserUseCase{validateIfUserExistsRepository, hasher, createUserRepository}
}

func (useCase *DbCreateUserUseCase) Execute(params usecases.CreateUserParams) (*models.User, error) {
	hasUser, err := useCase.validateIfUserExistsRepository.HasUser(params.Email)
	if err != nil {
		return nil, err
	}
	if hasUser {
		return nil, errors.New("user already exists")
	}
	hashedPassword, err := useCase.hasher.Hash(params.Password)
	if err != nil {
		return nil, err
	}
	params.Password = hashedPassword
	user, err := useCase.createUserRepository.CreateUser(params.Username, params.Password, params.Email, params.Logo)
	if err != nil {
		return nil, err
	}
	return user, nil
}
