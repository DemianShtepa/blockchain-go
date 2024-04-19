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

func randomBlock(t *testing.T, headerHeight uint64) *block.Block {
	return block.NewBlock(randomHeader(t, headerHeight), nil)
}

func randomBlockWithSignature(t *testing.T, headerHeight uint64, privateKey internal.PrivateKey) *block.Block {
	b := block.NewBlock(randomHeader(t, headerHeight), nil)
	_, err := b.Hash(&binary.BlockEncoder{})

	assert.Nil(t, err)
	assert.Nil(t, b.Sign(privateKey))

	return b
}

func randomPrivateKey(t *testing.T) internal.PrivateKey {
	privateKey, err := internal.NewPrivateKey()
	assert.Nil(t, err)

	return privateKey
}

func TestBlock_Hash(t *testing.T) {
	b := randomBlock(t, uint64(10))
	encoder := binary.BlockEncoder{}

	hash, err := b.Hash(&encoder)

	assert.Nil(t, err)
	assert.False(t, hash.IsEmpty())
}

func TestBlock_Sign_FailWithNoHash(t *testing.T) {
	b := randomBlock(t, uint64(10))
	privateKey := randomPrivateKey(t)

	assert.NotNil(t, b.Sign(privateKey))
}

func TestBlock_SignVerify(t *testing.T) {
	b := randomBlock(t, uint64(10))
	privateKey := randomPrivateKey(t)
	_, err := b.Hash(&binary.BlockEncoder{})

	assert.Nil(t, err)
	assert.Nil(t, b.Sign(privateKey))
	assert.True(t, b.Verify())
}
