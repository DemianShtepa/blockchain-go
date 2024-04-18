package block_test

import (
	"crypto/rand"
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/DemianShtepa/blockchain-go/internal/encode/binary"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomHash(t *testing.T) (internal.Hash, error) {
	value := make([]byte, internal.HashLength)

	_, err := rand.Read(value)
	assert.Nil(t, err)

	return internal.HashFromBytes(value)
}

func randomHeader(t *testing.T, height uint64) *block.Header {
	hash, err := randomHash(t)
	assert.Nil(t, err)

	return block.NewHeader(1, hash, time.Now().UnixNano(), height)
}

func randomBlock(t *testing.T) *block.Block {
	return block.NewBlock(randomHeader(t, 10), nil)
}

func randomPrivateKey(t *testing.T) internal.PrivateKey {
	privateKey, err := internal.NewPrivateKey()
	assert.Nil(t, err)

	return privateKey
}

func TestBlock_Hash(t *testing.T) {
	b := randomBlock(t)
	encoder := binary.BlockEncoder{}

	hash, err := b.Hash(&encoder)
	assert.Nil(t, err)
	assert.False(t, hash.IsEmpty())
}

func TestBlock_SignVerify(t *testing.T) {
	b := randomBlock(t)
	privateKey := randomPrivateKey(t)

	assert.Nil(t, b.Sign(privateKey))
	assert.True(t, b.Verify())
}
