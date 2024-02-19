package usecases

import (
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/asaskevich/govalidator"
)

type CreateUserParams struct {
	Username string `valid:"stringlength(5|50),optional"`
	Password string `valid:"required"`
	Email    string `valid:"email,required"`
	Logo     string `valid:"url,optional"`
}
type CreateUserUseCase interface {
	Execute(params CreateUserParams) (*models.User, error)
}

func (user *CreateUserParams) Validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}
