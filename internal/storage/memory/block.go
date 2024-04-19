package memory

import "github.com/DemianShtepa/blockchain-go/internal/block"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Put(b *block.Block) error {
	return nil
}
