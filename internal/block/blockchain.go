package block

import (
	"errors"
	"fmt"
	"sync"
)

type Storage interface {
	Put(*Block) error
}

type Blockchain struct {
	storage Storage
	lock    sync.RWMutex
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
	b.lock.Lock()
	defer b.lock.Unlock()

	if err := b.validateBlock(block); err != nil {
		return err
	}

	return b.addBlockWithoutValidation(block)
}

func (b *Blockchain) Height() uint64 {
	b.lock.RLock()
	defer b.lock.RUnlock()

	return b.unsafeHeight()
}

func (b *Blockchain) unsafeHeight() uint64 {
	return uint64(len(b.headers) - 1)
}

func (b *Blockchain) HasBlock(height uint64) bool {
	return b.Height() >= height
}

func (b *Blockchain) unsafeHasBlock(height uint64) bool {
	return b.unsafeHeight() >= height
}

func (b *Blockchain) LastHeader() *Header {
	b.lock.RLock()
	defer b.lock.RUnlock()

	return b.unsafeLastHeader()
}

func (b *Blockchain) unsafeLastHeader() *Header {
	return b.headers[len(b.headers)-1]
}

func (b *Blockchain) addBlockWithoutValidation(block *Block) error {
	b.headers = append(b.headers, block.Head)

	return b.storage.Put(block)
}

func (b *Blockchain) validateBlock(block *Block) error {
	blockHeight := block.Head.Height
	if b.unsafeHasBlock(blockHeight) {
		return fmt.Errorf("blockchain already contains block (%d)", block.Head.Height)
	}

	if blockHeight != b.unsafeHeight()+1 {
		return errors.New("inconsequent height of the block")
	}

	previousHeader := b.unsafeLastHeader()
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
