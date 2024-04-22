package block_test

import (
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/DemianShtepa/blockchain-go/internal/encode/binary"
	"github.com/DemianShtepa/blockchain-go/internal/storage/memory"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestBlockchain_AddBlock(t *testing.T) {
	privateKey := randomPrivateKey(t)
	headerEncoder := binary.HeaderEncoder{}
	genesisBlockHeader := block.NewHeader(1, internal.Hash{}, time.Now().UnixNano(), uint64(0), &headerEncoder)
	genesisBlock := block.NewBlock(genesisBlockHeader, nil)
	_, err := genesisBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, genesisBlock.Sign(privateKey))
	blockchain, err := block.NewBlockchain(memory.NewStorage(), genesisBlock)
	assert.Nil(t, err)

	genesisBlockHash, err := genesisBlock.Hash()
	assert.Nil(t, err)
	header := block.NewHeader(1, genesisBlockHash, time.Now().UnixNano(), uint64(1), &headerEncoder)
	newBlock := block.NewBlock(header, nil)
	_, err = newBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, newBlock.Sign(privateKey))
	assert.Nil(t, blockchain.AddBlock(newBlock))

	assert.Equal(t, uint64(1), blockchain.Height())
}

func TestBlockchain_AddBlock_FailWithExistingBlock(t *testing.T) {
	privateKey := randomPrivateKey(t)
	headerEncoder := binary.HeaderEncoder{}
	genesisBlockHeader := block.NewHeader(1, internal.Hash{}, time.Now().UnixNano(), uint64(0), &headerEncoder)
	genesisBlock := block.NewBlock(genesisBlockHeader, nil)
	_, err := genesisBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, genesisBlock.Sign(privateKey))
	blockchain, err := block.NewBlockchain(memory.NewStorage(), genesisBlock)
	assert.Nil(t, err)

	genesisBlockHash, err := genesisBlock.Hash()
	assert.Nil(t, err)
	header := block.NewHeader(1, genesisBlockHash, time.Now().UnixNano(), uint64(0), &headerEncoder)
	newBlock := block.NewBlock(header, nil)
	_, err = newBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, newBlock.Sign(privateKey))
	assert.NotNil(t, blockchain.AddBlock(newBlock))

	assert.Equal(t, uint64(0), blockchain.Height())
}

func TestBlockchain_AddBlock_FailWithInconsequentHeight(t *testing.T) {
	privateKey := randomPrivateKey(t)
	headerEncoder := binary.HeaderEncoder{}
	genesisBlockHeader := block.NewHeader(1, internal.Hash{}, time.Now().UnixNano(), uint64(0), &headerEncoder)
	genesisBlock := block.NewBlock(genesisBlockHeader, nil)
	_, err := genesisBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, genesisBlock.Sign(privateKey))
	blockchain, err := block.NewBlockchain(memory.NewStorage(), genesisBlock)
	assert.Nil(t, err)

	genesisBlockHash, err := genesisBlock.Hash()
	assert.Nil(t, err)
	header := block.NewHeader(1, genesisBlockHash, time.Now().UnixNano(), uint64(2), &headerEncoder)
	newBlock := block.NewBlock(header, nil)
	_, err = newBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, newBlock.Sign(privateKey))
	assert.NotNil(t, blockchain.AddBlock(newBlock))

	assert.Equal(t, uint64(0), blockchain.Height())
}

func TestBlockchain_AddBlock_FailWithInvalidPreviousBlockHash(t *testing.T) {
	privateKey := randomPrivateKey(t)
	headerEncoder := binary.HeaderEncoder{}
	genesisBlockHeader := block.NewHeader(1, internal.Hash{}, time.Now().UnixNano(), uint64(0), &headerEncoder)
	genesisBlock := block.NewBlock(genesisBlockHeader, nil)
	_, err := genesisBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, genesisBlock.Sign(privateKey))
	blockchain, err := block.NewBlockchain(memory.NewStorage(), genesisBlock)
	assert.Nil(t, err)

	genesisBlockHash := randomHash(t)
	assert.Nil(t, err)
	header := block.NewHeader(1, genesisBlockHash, time.Now().UnixNano(), uint64(1), &headerEncoder)
	newBlock := block.NewBlock(header, nil)
	_, err = newBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, newBlock.Sign(privateKey))
	assert.NotNil(t, blockchain.AddBlock(newBlock))

	assert.Equal(t, uint64(0), blockchain.Height())
}

func TestBlockchain_Height(t *testing.T) {
	privateKey := randomPrivateKey(t)
	headerEncoder := binary.HeaderEncoder{}
	genesisBlockHeader := block.NewHeader(1, internal.Hash{}, time.Now().UnixNano(), uint64(0), &headerEncoder)
	genesisBlock := block.NewBlock(genesisBlockHeader, nil)
	_, err := genesisBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, genesisBlock.Sign(privateKey))
	blockchain, err := block.NewBlockchain(memory.NewStorage(), genesisBlock)
	assert.Nil(t, err)

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
	headerEncoder := binary.HeaderEncoder{}
	genesisBlockHeader := block.NewHeader(1, internal.Hash{}, time.Now().UnixNano(), uint64(0), &headerEncoder)
	genesisBlock := block.NewBlock(genesisBlockHeader, nil)
	_, err := genesisBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, genesisBlock.Sign(privateKey))
	blockchain, err := block.NewBlockchain(memory.NewStorage(), genesisBlock)
	assert.Nil(t, err)

	genesisBlockHash, err := genesisBlock.Hash()
	assert.Nil(t, err)
	header := block.NewHeader(1, genesisBlockHash, time.Now().UnixNano(), uint64(1), &headerEncoder)
	newBlock := block.NewBlock(header, nil)
	_, err = newBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, newBlock.Sign(privateKey))
	assert.Nil(t, blockchain.AddBlock(newBlock))

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.hasBlock, blockchain.HasBlock(c.height))
		})
	}
}

func TestBlockchain_ConcurrentAccess(t *testing.T) {
	privateKey := randomPrivateKey(t)
	headerEncoder := binary.HeaderEncoder{}
	genesisBlockHeader := block.NewHeader(1, internal.Hash{}, time.Now().UnixNano(), uint64(0), &headerEncoder)
	genesisBlock := block.NewBlock(genesisBlockHeader, nil)
	_, err := genesisBlock.Hash()
	assert.Nil(t, err)
	assert.Nil(t, genesisBlock.Sign(privateKey))
	blockchain, err := block.NewBlockchain(memory.NewStorage(), genesisBlock)
	assert.Nil(t, err)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()

		genesisBlockHash, err := genesisBlock.Hash()
		assert.Nil(t, err)
		header := block.NewHeader(1, genesisBlockHash, time.Now().UnixNano(), uint64(1), &headerEncoder)
		newBlock := block.NewBlock(header, nil)
		_, err = newBlock.Hash()
		assert.Nil(t, err)
		assert.Nil(t, newBlock.Sign(privateKey))
		assert.Nil(t, blockchain.AddBlock(newBlock))
	}()

	go func() {
		defer wg.Done()

		assert.NotNil(t, blockchain.Height())
	}()

	go func() {
		defer wg.Done()

		assert.NotNil(t, blockchain.LastHeader())
	}()

	wg.Wait()

	assert.Equal(t, uint64(1), blockchain.Height())
}
