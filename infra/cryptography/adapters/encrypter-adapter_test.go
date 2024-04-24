package adapters_test

import (
	"testing"
	"time"

	"github.com/RafaelCava/chat-auth-go/domain/models"
	"github.com/RafaelCava/chat-auth-go/infra/cryptography/adapters"
	"github.com/stretchr/testify/assert"
)

func TestEncrypterAdapterEncryptSuccess(t *testing.T) {
	sut := adapters.NewEncrypterAdapter("any_secret", "any_issuer")
	token, err := sut.Encrypt(&models.Claims{}, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, token)
	assert.NotEmpty(t, token)
}

func TestEncrypterAdapterDecryptSuccess(t *testing.T) {
	sut := adapters.NewEncrypterAdapter("any_secret", "any_issuer")
	token, err := sut.Encrypt(&models.Claims{}, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := sut.Decrypt(token)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, claims)
	assert.Equal(t, "any_issuer", claims.Issuer)
}

func TestEncrypterAdapterDecryptErrorTime(t *testing.T) {
	sut := adapters.NewEncrypterAdapter("any_secret", "any_issuer")
	token, err := sut.Encrypt(&models.Claims{}, time.Microsecond)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)
	claims, err := sut.Decrypt(token)
	assert.Nil(t, claims)
	assert.Error(t, err)
	assert.Equal(t, "token is expired by 1s", err.Error())
}

func TestEncrypterAdapterDecryptErrorIssuer(t *testing.T) {
	sutFake := adapters.NewEncrypterAdapter("any_secret", "any_issuer_fake")
	tokenFake, errFake := sutFake.Encrypt(&models.Claims{}, time.Microsecond)
	if errFake != nil {
		t.Fatal(errFake)
	}
	sut := adapters.NewEncrypterAdapter("any_secret", "any_issuer")
	claims, err := sut.Decrypt(tokenFake)
	assert.Nil(t, claims)
	assert.Error(t, err)
	assert.Equal(t, "invalid token", err.Error())
}

func TestEncrypterAdapterDecryptErrorAssign(t *testing.T) {
	sut := adapters.NewEncrypterAdapter("any_secret", "any_issuer")
	claims, err := sut.Decrypt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	assert.Nil(t, claims)
	assert.Error(t, err)
	assert.Equal(t, "signature is invalid", err.Error())
}

func TestEncrypterAdapterDecryptErrorClaimsConversion(t *testing.T) {
	sut := adapters.NewEncrypterAdapter("any_secret", "any_issuer")
	claims, err := sut.Decrypt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.JOxTQsNDINOUToIgWdZYQWXG0T2zp8ItXNzO9mvfPqs")
	assert.Nil(t, claims)
	assert.Error(t, err)
	assert.Equal(t, "expired credential", err.Error())
}
