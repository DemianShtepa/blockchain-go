package internal

import (
	"bytes"
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomHash() (Hash, error) {
	value := make([]byte, hashLength)

	if _, err := rand.Read(value); err != nil {
		return Hash{}, err
	}

	return HashFromBytes(value)
}

func TestHeader_EncodeBinary_DecodeBinary(t *testing.T) {
	hash, err := randomHash()
	assert.Nil(t, err)

	header := NewHeader(1, hash, time.Now().UnixNano(), 5, 58342193)

	buf := bytes.Buffer{}
	assert.Nil(t, header.EncodeBinary(&buf))

	var decodedHeader Header
	assert.Nil(t, decodedHeader.DecodeBinary(&buf))

	assert.Equal(t, *header, decodedHeader)
}

func TestBlock_Hash(t *testing.T) {
	hash, err := randomHash()
	assert.Nil(t, err)

	block := NewBlock(NewHeader(1, hash, time.Now().UnixNano(), 5, 58342193), nil)

	blockHash, err := block.Hash()
	assert.Nil(t, err)
	assert.False(t, blockHash.IsEmpty())
}

func TestBlock_EncodeBinary_DecodeBinary(t *testing.T) {
	hash, err := randomHash()
	assert.Nil(t, err)

	block := NewBlock(NewHeader(1, hash, time.Now().UnixNano(), 5, 58342193), nil)

	buf := bytes.Buffer{}
	assert.Nil(t, block.EncodeBinary(&buf))

	decodedBlock := NewBlock(&Header{}, nil)
	assert.Nil(t, decodedBlock.DecodeBinary(&buf))

	assert.Equal(t, *block.Head, *decodedBlock.Head)
	assert.Equal(t, block.Transactions, decodedBlock.Transactions)
}
