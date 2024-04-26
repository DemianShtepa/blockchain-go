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

func randomHash(t *testing.T) internal.Hash {
	value := make([]byte, internal.HashLength)

	_, err := rand.Read(value)
	assert.Nil(t, err)

	return internal.HashFromBytes(value)
}

func TestHeader_EncodeDecode(t *testing.T) {
	var buf bytes.Buffer

	hash := randomHash(t)
	encoder := gob.HeaderEncoder{}
	header := block.NewHeader(1, hash, time.Now().UnixNano(), 10, &encoder)

	assert.Nil(t, encoder.Encode(&buf, header))

	decodedHeader := block.NewHeader(1, hash, time.Now().Add(time.Hour).UnixNano(), 1, &encoder)
	decoder := gob.HeaderDecoder{}

	assert.Nil(t, decoder.Decode(&buf, decodedHeader))

	assert.Equal(t, *header, *decodedHeader)
}
