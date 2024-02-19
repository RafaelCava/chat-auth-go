package protocols

import (
	"time"

	"github.com/RafaelCava/chat-auth-go/domain/models"
)

type Encrypter interface {
	Encrypt(claims *models.Claims, ttl time.Duration) (string, error)
}
