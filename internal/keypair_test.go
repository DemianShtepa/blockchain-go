package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignature_Verify_Success(t *testing.T) {
	privateKey, _ := NewPrivateKey()
	data := []byte("Test")

	signature, err := privateKey.Sign(data)

	assert.Nil(t, err)
	assert.True(t, signature.Verify(privateKey.PublicKey(), data))
}

func TestSignature_Verify_FailWithDifferentPublicKey(t *testing.T) {
	privateKey, _ := NewPrivateKey()
	differentPrivateKey, _ := NewPrivateKey()
	data := []byte("Test")

	signature, err := privateKey.Sign(data)

	assert.Nil(t, err)
	assert.False(t, signature.Verify(differentPrivateKey.PublicKey(), data))
}

func TestSignature_Verify_FailWithDifferentData(t *testing.T) {
	privateKey, _ := NewPrivateKey()
	data := []byte("Test")

	signature, err := privateKey.Sign(data)

	assert.Nil(t, err)
	assert.False(t, signature.Verify(privateKey.PublicKey(), []byte("Value")))
}
