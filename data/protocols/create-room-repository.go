package protocols

import "github.com/RafaelCava/chat-auth-go/domain/models"

type CreateRoomRepositoryParams struct {
	Name    string `json:"name"`
	OwnerId string `json:"owner_id"`
}

type CreateRoomRepository interface {
	Create(params CreateRoomRepositoryParams) (*models.Room, error)
}
