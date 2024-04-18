package block

import (
	"github.com/DemianShtepa/blockchain-go/internal"
)

type Block struct {
	Head         *Header
	Transactions Transactions

	headerHash internal.Hash
}

func NewBlock(head *Header, transactions []Transaction) *Block {
	return &Block{Head: head, Transactions: transactions}
}
