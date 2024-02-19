package repositories

import (
	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"gorm.io/gorm"
)

type FindUserByEmailPostgresRepository struct {
	db *gorm.DB
}

func NewFindUserByEmailPostgresRepository(db *gorm.DB) protocols.FindUserByEmailRepository {
	return &FindUserByEmailPostgresRepository{db}
}

func (repository *FindUserByEmailPostgresRepository) FindByEmail(email string, projection []string) (*models.User, error) {
	var user models.User
	result := repository.db.Select(projection).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
