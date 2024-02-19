package adapters

import (
	"errors"
	"time"

	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/dgrijalva/jwt-go"
)

type EncrypterAdapter struct {
	secretKey string
	issuer    string
}

type Encrypter interface {
	protocols.Encrypter
	protocols.Decrypter
}

func NewEncrypterAdapter(secretKey string, issuer string) Encrypter {
	return &EncrypterAdapter{
		secretKey,
		issuer,
	}
}

func (adapter *EncrypterAdapter) Encrypt(claims *models.Claims, duration time.Duration) (string, error) {
	tokenClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(duration).Unix(),
		Issuer:    adapter.issuer,
	}
	claims.StandardClaims = tokenClaims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, err := token.SignedString([]byte(adapter.secretKey))
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}

func (adapter *EncrypterAdapter) Decrypt(token string) (*models.Claims, error) {
	claims := &models.Claims{}
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(adapter.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*models.Claims); ok && tokenClaims.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("expired credential")
		}
		if claims.Issuer != adapter.issuer {
			return nil, errors.New("invalid token")
		}
		return claims, nil
	} else {
		return nil, errors.New("invalid credentials")
	}
}
