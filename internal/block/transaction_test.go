package block_test

import (
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_Sign_Verify(t *testing.T) {
	privateKey, _ := internal.NewPrivateKey()
	transaction := block.Transaction{
		Data: []byte("Test"),
	}

	assert.Nil(t, transaction.Sign(privateKey))
	assert.True(t, transaction.Verify())
}

func TestTransaction_Verify_FailWithDifferentData(t *testing.T) {
	privateKey, _ := internal.NewPrivateKey()
	transaction := block.Transaction{
		Data: []byte("Test"),
	}

	assert.Nil(t, transaction.Sign(privateKey))

	transaction.Data = []byte("Value")
	assert.False(t, transaction.Verify())
}
