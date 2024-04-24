package repositories

import (
	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"gorm.io/gorm"
)

type CreateRoomPostgresRepository struct {
	db *gorm.DB
}

func NewCreateRoomPostgresRepository(db *gorm.DB) protocols.CreateRoomRepository {
	return &CreateRoomPostgresRepository{db}
}

func (repository *CreateRoomPostgresRepository) Create(params protocols.CreateRoomRepositoryParams) (*models.Room, error) {
	room, err := models.NewRoom(params.Name, params.OwnerId)
	if err != nil {
		return nil, err
	}
	if err := repository.db.Create(&room).Error; err != nil {
		return nil, err
	}
	return &models.Room{
		ID:        room.ID,
		Name:      room.Name,
		OwnerId:   room.OwnerId,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}, nil
}
