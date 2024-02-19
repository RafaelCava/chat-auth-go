package repositories

import (
	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"gorm.io/gorm"
)

type CreateUserPostgresRepository struct {
	db *gorm.DB
}

func NewCreateUserPostgresRepository(db *gorm.DB) protocols.CreateUserRepository {
	return &CreateUserPostgresRepository{db}
}

func (repository *CreateUserPostgresRepository) CreateUser(username, password, email, logo string) (*models.User, error) {
	user, err := models.NewUser(username, password, email, logo)
	if err != nil {
		return nil, err
	}
	if err := repository.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &models.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Logo:      user.Logo,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
