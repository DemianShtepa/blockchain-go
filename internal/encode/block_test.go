package encode

import (
	"bytes"
	"crypto/rand"
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomHash() (internal.Hash, error) {
	value := make([]byte, internal.HashLength)

	if _, err := rand.Read(value); err != nil {
		return internal.Hash{}, err
	}

	return internal.HashFromBytes(value)
}

func randomBlock(t *testing.T) *block.Block {
	hash, err := randomHash()
	assert.Nil(t, err)

	header := block.NewHeader(1, hash, time.Now().UnixNano(), 10)

	return block.NewBlock(header, nil)
}

func TestBlock_EncodeDecode(t *testing.T) {
	var buf bytes.Buffer

	b := randomBlock(t)
	encoder := NewBlockEncoder(&buf, NewHeaderEncoder(&buf), NewTransactionEncoder(&buf))

	assert.Nil(t, encoder.Encode(b))

	decodedBlock := block.Block{
		Head: &block.Header{},
	}
	decoder := NewBlockDecoder(&buf, NewHeaderDecoder(&buf), NewTransactionDecoder(&buf))

	assert.Nil(t, decoder.Decode(&decodedBlock))

	assert.Equal(t, *b, decodedBlock)
}
