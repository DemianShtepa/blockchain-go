package internal_test

import (
	"bytes"
	"crypto/rand"
	"github.com/DemianShtepa/blockchain-go/internal"
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
	hash := internal.HashFromBytes(randomBytes(t))

	assert.False(t, hash.IsEmpty())
}

func TestHashFromReader(t *testing.T) {
	var buf bytes.Buffer
	buf.Write(randomBytes(t))

	hash, err := internal.HashFromReader(&buf)

	assert.False(t, hash.IsEmpty())
	assert.Nil(t, err)
}
