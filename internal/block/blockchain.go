package block

import (
	"errors"
	"fmt"
)

type Storage interface {
	Put(*Block) error
}

type Blockchain struct {
	storage Storage
	headers []*Header
}

func NewBlockchain(storage Storage, genesisBlock *Block) (*Blockchain, error) {
	blockChain := &Blockchain{
		storage: storage,
		headers: []*Header{},
	}

	err := blockChain.addBlockWithoutValidation(genesisBlock)

	return blockChain, err
}

func (b *Blockchain) AddBlock(block *Block) error {
	if err := b.validateBlock(block); err != nil {
		return err
	}

	return b.addBlockWithoutValidation(block)
}

func (b *Blockchain) Height() uint64 {
	return uint64(len(b.headers) - 1)
}

func (b *Blockchain) HasBlock(height uint64) bool {
	return b.Height() >= height
}

func (b *Blockchain) addBlockWithoutValidation(block *Block) error {
	b.headers = append(b.headers, block.Head)

	return b.storage.Put(block)
}

func (b *Blockchain) validateBlock(block *Block) error {
	if b.HasBlock(block.Head.Height) {
		return fmt.Errorf("blockchain already contains block (%d)", block.Head.Height)
	}

	if !block.Verify() {
		return errors.New("invalid signature for block")
	}

	return nil
}
