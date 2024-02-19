package usecases

import (
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/asaskevich/govalidator"
)

type AuthRequest struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

type AuthUseCase interface {
	Execute(params AuthRequest) (*models.TokenPair, error)
}

func (authRequest *AuthRequest) Validate() error {
	_, err := govalidator.ValidateStruct(authRequest)

	if err != nil {
		return err
	}
	return nil
}
