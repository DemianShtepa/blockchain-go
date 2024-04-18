package gob_test

import (
	"bytes"
	"crypto/rand"
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/DemianShtepa/blockchain-go/internal/encode/gob"
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

func TestHeader_EncodeDecode(t *testing.T) {
	var buf bytes.Buffer

	hash, _ := randomHash(t)
	header := block.NewHeader(1, hash, time.Now().UnixNano(), 10)
	encoder := gob.HeaderEncoder{}

	assert.Nil(t, encoder.Encode(&buf, header))

	var decodedHeader block.Header
	decoder := gob.HeaderDecoder{}

	assert.Nil(t, decoder.Decode(&buf, &decodedHeader))

	assert.Equal(t, *header, decodedHeader)
}
