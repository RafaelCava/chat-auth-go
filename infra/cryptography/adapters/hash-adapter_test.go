package adapters_test

import (
	"testing"

	"github.com/RafaelCava/chat-auth-go/infra/cryptography/adapters"
	"github.com/stretchr/testify/assert"
)

func TestHasherAdapter(t *testing.T) {
	hasherAdapter := adapters.NewHasherAdapter()
	hashedValue, err := hasherAdapter.Hash("any_value")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, hashedValue)
}

func TestHasherAdapterComparer(t *testing.T) {
	hasherAdapter := adapters.NewHasherAdapter()
	hashedValue, err := hasherAdapter.Hash("any_value")
	if err != nil {
		t.Fatal(err)
	}
	err = hasherAdapter.Compare("any_value", hashedValue)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
}

func TestHasherAdapterComparerError(t *testing.T) {
	hasherAdapter := adapters.NewHasherAdapter()
	err := hasherAdapter.Compare("any_value", "any_hashed_value")
	assert.Error(t, err)
}
