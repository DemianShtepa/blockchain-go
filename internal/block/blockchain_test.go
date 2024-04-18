package block

import (
	"crypto/rand"
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomHash() (internal.Hash, error) {
	value := make([]byte, internal.HashLength)

	if _, err := rand.Read(value); err != nil {
		return internal.Hash{}, err
	}

	return internal.HashFromBytes(value)
}

func randomBlock(t *testing.T) *Block {
	hash, err := randomHash()
	assert.Nil(t, err)

	header := NewHeader(1, hash, time.Now().UnixNano(), 10)

	return NewBlock(header, nil)
}

func newBlockchain(t *testing.T) *Blockchain {
	return NewBlockchain(nil, randomBlock(t))
}

func TestBlockchain_Height(t *testing.T) {
	blockchain := newBlockchain(t)

	assert.Equal(t, uint64(0), blockchain.Height())
}
