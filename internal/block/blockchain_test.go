package block_test

import (
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/DemianShtepa/blockchain-go/internal/storage/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newBlockchain(t *testing.T, headerHeight uint64, privateKey internal.PrivateKey) *block.Blockchain {
	blockchain, err := block.NewBlockchain(memory.NewStorage(), randomBlockWithSignature(t, headerHeight, privateKey))
	assert.Nil(t, err)

	return blockchain
}

func TestBlockchain_AddBlock(t *testing.T) {
	privateKey := randomPrivateKey(t)
	blockchain := newBlockchain(t, 0, privateKey)

	assert.Nil(t, blockchain.AddBlock(randomBlockWithSignature(t, 1, privateKey)))

	assert.Equal(t, uint64(1), blockchain.Height())
}

func TestBlockchain_AddBlock_FailWithInvalidHeight(t *testing.T) {
	privateKey := randomPrivateKey(t)
	blockchain := newBlockchain(t, 0, privateKey)

	assert.NotNil(t, blockchain.AddBlock(randomBlockWithSignature(t, 0, privateKey)))

	assert.Equal(t, uint64(0), blockchain.Height())
}

func TestBlockchain_Height(t *testing.T) {
	blockchain := newBlockchain(t, 0, randomPrivateKey(t))

	assert.Equal(t, uint64(0), blockchain.Height())
}

func TestBlockchain_HasBlock(t *testing.T) {
	cases := []struct {
		title    string
		height   uint64
		hasBlock bool
	}{
		{
			title:    "genesis block height is equal",
			height:   0,
			hasBlock: true,
		},
		{
			title:    "block header height is less",
			height:   1,
			hasBlock: true,
		},
		{
			title:    "block header height is higher",
			height:   5,
			hasBlock: false,
		},
	}

	privateKey := randomPrivateKey(t)
	blockchain := newBlockchain(t, 0, privateKey)
	for i := range cases {
		assert.Nil(t, blockchain.AddBlock(randomBlockWithSignature(t, uint64(i+1), privateKey)))
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.hasBlock, blockchain.HasBlock(c.height))
		})
	}
}
