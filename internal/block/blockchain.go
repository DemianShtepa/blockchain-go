package block

type Storage interface {
}

type Blockchain struct {
	storage Storage
	headers []*Header
}

func NewBlockchain(storage Storage, genesisBlock *Block) *Blockchain {
	return &Blockchain{
		storage: storage,
		headers: []*Header{genesisBlock.Head},
	}
}

func (b *Blockchain) Height() uint64 {
	return uint64(len(b.headers) - 1)
}
