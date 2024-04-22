package binary_test

import (
	"bytes"
	"crypto/rand"
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/DemianShtepa/blockchain-go/internal/encode/binary"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomHash(t *testing.T) (internal.Hash, error) {
	value := make([]byte, internal.HashLength)

	_, err := rand.Read(value)
	assert.Nil(t, err)

	return internal.HashFromBytes(value)
}

func randomHeader(t *testing.T) *block.Header {
	hash, err := randomHash(t)
	assert.Nil(t, err)

	return block.NewHeader(1, hash, time.Now().UnixNano(), 10, &binary.HeaderEncoder{})
}

func randomBlock(t *testing.T) *block.Block {
	return block.NewBlock(randomHeader(t), nil)
}

func TestBlock_EncodeDecode(t *testing.T) {
	var buf bytes.Buffer

	b := randomBlock(t)
	encoder := binary.BlockEncoder{}

	assert.Nil(t, encoder.Encode(&buf, b))

	decodedBlock := block.Block{
		Head: randomHeader(t),
	}
	decoder := binary.BlockDecoder{}

	assert.Nil(t, decoder.Decode(&buf, &decodedBlock))

	assert.Equal(t, *b, decodedBlock)
}
