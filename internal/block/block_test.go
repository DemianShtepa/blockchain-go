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

func randomHash(t *testing.T) internal.Hash {
	value := make([]byte, internal.HashLength)

	_, err := rand.Read(value)
	assert.Nil(t, err)

	hash, err := internal.HashFromBytes(value)
	assert.Nil(t, err)

	return hash
}

func randomHeader(t *testing.T, height uint64) *block.Header {
	hash := randomHash(t)

	return block.NewHeader(1, hash, time.Now().UnixNano(), height, &binary.HeaderEncoder{})
}

func randomBlock(t *testing.T, headerHeight uint64) *block.Block {
	privateKey := randomPrivateKey(t)
	transaction := block.Transaction{Data: []byte("Test")}
	assert.Nil(t, transaction.Sign(privateKey))

	return block.NewBlock(
		randomHeader(t, headerHeight),
		block.Transactions{transaction},
	)
}

func randomPrivateKey(t *testing.T) internal.PrivateKey {
	privateKey, err := internal.NewPrivateKey()
	assert.Nil(t, err)

	return privateKey
}

func TestBlock_Hash(t *testing.T) {
	b := randomBlock(t, uint64(10))

	hash, err := b.Hash()

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
	_, err := b.Hash()

	assert.Nil(t, err)
	assert.Nil(t, b.Sign(privateKey))
	assert.True(t, b.Verify())
}
