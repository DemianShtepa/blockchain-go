package network_test

import (
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/DemianShtepa/blockchain-go/internal/network"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPool_Add(t *testing.T) {
	transaction := block.NewTransaction([]byte("Test"))
	pool := network.NewPool()

	assert.Nil(t, pool.Add(transaction))
	assert.Equal(t, 1, pool.Len())
}

func TestPool_Add_FailWithTransactionDuplication(t *testing.T) {
	transaction := block.NewTransaction([]byte("Test"))
	pool := network.NewPool()

	assert.Nil(t, pool.Add(transaction))
	assert.ErrorIs(t, pool.Add(transaction), network.ErrTransactionAlreadyExists)
	assert.Equal(t, 1, pool.Len())
}

func TestPool_Flush(t *testing.T) {
	transaction := block.NewTransaction([]byte("Test"))
	pool := network.NewPool()

	assert.Nil(t, pool.Add(transaction))

	pool.Flush()

	assert.Equal(t, 0, pool.Len())
}
