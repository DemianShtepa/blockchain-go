package internal

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func randomBytes(t *testing.T) []byte {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	assert.Nil(t, err)

	return b
}

func TestHashFromBytes(t *testing.T) {
	hash, err := HashFromBytes(randomBytes(t))
	assert.False(t, hash.IsEmpty())
	assert.Nil(t, err)
}
