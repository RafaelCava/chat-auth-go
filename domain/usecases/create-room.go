package usecases

import (
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/asaskevich/govalidator"
)

type CreateRoomParams struct {
	Name    string `valid:"stringlength(5|50),required"`
	OwnerId string `valid:"uuid,required"`
}

type CreateRoomUseCase interface {
	Execute(params CreateRoomParams) (*models.Room, error)
}

func (data *CreateRoomParams) Validate() error {
	_, err := govalidator.ValidateStruct(data)

	if err != nil {
		return err
	}

	return nil
}
