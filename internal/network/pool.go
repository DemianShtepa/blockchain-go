package network

import (
	"errors"
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
)

var ErrTransactionAlreadyExists = errors.New("transaction already exists in the pool")

type Pool struct {
	transactions map[internal.Hash]*block.Transaction
}

func NewPool() *Pool {
	return &Pool{
		transactions: make(map[internal.Hash]*block.Transaction),
	}
}

func (p *Pool) Add(transaction *block.Transaction) error {
	hash := transaction.Hash()

	if _, ok := p.transactions[hash]; ok {
		return ErrTransactionAlreadyExists
	}

	p.transactions[hash] = transaction

	return nil
}

func (p *Pool) Len() int {
	return len(p.transactions)
}

func (p *Pool) Flush() {
	p.transactions = make(map[internal.Hash]*block.Transaction)
}
