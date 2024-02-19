package repositories

import (
	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"gorm.io/gorm"
)

type ValidateIfUserExistsPostgresRepository struct {
	db *gorm.DB
}

func NewValidateIfUserExistsPostgresRepository(db *gorm.DB) protocols.ValidateIfUserExistsRepository {
	return &ValidateIfUserExistsPostgresRepository{db}
}

func (repository *ValidateIfUserExistsPostgresRepository) HasUser(email string) (bool, error) {
	var user models.User
	if err := repository.db.Select("email").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
