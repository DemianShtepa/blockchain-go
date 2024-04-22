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

func (b *Blockchain) GetLastHeader() *Header {
	return b.headers[len(b.headers)-1]
}

func (b *Blockchain) GetHeader(height uint64) (*Header, error) {
	if !b.HasBlock(height) {
		return nil, errors.New("block doesn't exist")
	}

	return b.headers[height], nil
}

func (b *Blockchain) addBlockWithoutValidation(block *Block) error {
	b.headers = append(b.headers, block.Head)

	return b.storage.Put(block)
}

func (b *Blockchain) validateBlock(block *Block) error {
	blockHeight := block.Head.Height
	if b.HasBlock(blockHeight) {
		return fmt.Errorf("blockchain already contains block (%d)", block.Head.Height)
	}

	if blockHeight != b.Height()+1 {
		return errors.New("inconsequent height of the block")
	}

	previousHeader := b.GetLastHeader()
	previousHeaderHash, err := previousHeader.Hash()
	if err != nil {
		return err
	}
	if previousHeaderHash != block.Head.PreviousBlockHash {
		return errors.New("invalid previous block hash provided")
	}

	if !block.Verify() {
		return errors.New("invalid signature of the block")
	}

	return nil
}
