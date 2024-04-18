package block

import (
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_Sign_Verify(t *testing.T) {
	privateKey, _ := internal.NewPrivateKey()
	transaction := Transaction{
		Data: []byte("Test"),
	}

	assert.Nil(t, transaction.Sign(privateKey))
	assert.True(t, transaction.Verify())
}

func TestTransaction_Verify_FailWithDifferentPublicKey(t *testing.T) {
	privateKey, _ := internal.NewPrivateKey()
	differentPrivateKey, _ := internal.NewPrivateKey()
	transaction := Transaction{
		Data: []byte("Test"),
	}

	assert.Nil(t, transaction.Sign(privateKey))

	transaction.PublicKey = differentPrivateKey.PublicKey()
	assert.False(t, transaction.Verify())
}

func TestTransaction_Verify_FailWithDifferentData(t *testing.T) {
	privateKey, _ := internal.NewPrivateKey()
	transaction := Transaction{
		Data: []byte("Test"),
	}

	assert.Nil(t, transaction.Sign(privateKey))

	transaction.Data = []byte("Value")
	assert.False(t, transaction.Verify())
}

func TestTransaction_Verify_FailWithDifferentSignature(t *testing.T) {
	privateKey, _ := internal.NewPrivateKey()
	transaction := Transaction{
		Data: []byte("Test"),
	}

	assert.Nil(t, transaction.Sign(privateKey))

	transaction.Signature = internal.Signature{}
	assert.False(t, transaction.Verify())
}
