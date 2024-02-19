package adapters

import (
	"github.com/RafaelCava/chat-auth-go/data/protocols"
	"golang.org/x/crypto/bcrypt"
)

type HasherAdapter struct{}

type Hasher interface {
	protocols.Hasher
	protocols.HasherCompare
}

func NewHasherAdapter() Hasher {
	return &HasherAdapter{}
}

func (adapter *HasherAdapter) Hash(value string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), 11)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (adapter *HasherAdapter) Compare(value string, hashedValue string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
}
