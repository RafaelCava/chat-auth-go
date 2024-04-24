package usecases

import (
	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/RafaelCava/chat-auth-go/domain/usecases"
)

type DbCreateRoomUseCase struct {
	CreateRoomRepository protocols.CreateRoomRepository
}

func NewDbCreateRoomUseCase(createRoomRepository protocols.CreateRoomRepository) usecases.CreateRoomUseCase {
	return &DbCreateRoomUseCase{createRoomRepository}
}

func (useCase *DbCreateRoomUseCase) Execute(params usecases.CreateRoomParams) (*models.Room, error) {
	room, err := useCase.CreateRoomRepository.Create(protocols.CreateRoomRepositoryParams{
		Name:    params.Name,
		OwnerId: params.OwnerId,
	})
	if err != nil {
		return nil, err
	}
	return room, nil
}
