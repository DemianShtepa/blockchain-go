package block_test

import (
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newBlockchain(t *testing.T) *block.Blockchain {
	return block.NewBlockchain(nil, randomBlock(t))
}

func TestBlockchain_Height(t *testing.T) {
	blockchain := newBlockchain(t)

	assert.Equal(t, uint64(0), blockchain.Height())
}
