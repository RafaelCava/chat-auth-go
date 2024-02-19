package usecases

import (
	"time"

	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/RafaelCava/chat-auth-go/domain/usecases"
)

type DbLogin struct {
	findUserByEmailRepository protocols.FindUserByEmailRepository
	hasherCompare             protocols.HasherCompare
	encrypter                 protocols.Encrypter
}

func NewDbLogin(findUserByEmailRepository protocols.FindUserByEmailRepository, hasherCompare protocols.HasherCompare, encrypter protocols.Encrypter) usecases.AuthUseCase {
	return &DbLogin{findUserByEmailRepository, hasherCompare, encrypter}
}

func (dbLogin *DbLogin) Execute(params usecases.AuthRequest) (*models.TokenPair, error) {
	user, err := dbLogin.findUserByEmailRepository.FindByEmail(params.Email, []string{"password", "id"})
	if err != nil {
		return nil, err
	}

	if err := dbLogin.hasherCompare.Compare(params.Password, user.Password); err != nil {
		return nil, err
	}

	accessToken, err := dbLogin.encrypter.Encrypt(&models.Claims{UserId: user.ID}, 3*time.Hour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := dbLogin.encrypter.Encrypt(&models.Claims{UserId: user.ID}, time.Hour*24*4)
	if err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
