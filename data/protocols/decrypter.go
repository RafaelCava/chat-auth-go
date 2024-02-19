package protocols

import "github.com/RafaelCava/chat-auth-go/domain/models"

type Decrypter interface {
	Decrypt(ciphertext string) (*models.Claims, error)
}
