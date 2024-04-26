package binary_test

import (
	"bytes"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/DemianShtepa/blockchain-go/internal/encode/binary"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHeader_EncodeDecode(t *testing.T) {
	var buf bytes.Buffer

	hash := randomHash(t)
	encoder := binary.HeaderEncoder{}
	header := block.NewHeader(1, hash, time.Now().UnixNano(), 10, &encoder)

	assert.Nil(t, encoder.Encode(&buf, header))

	decodedHeader := block.NewHeader(1, hash, time.Now().Add(time.Hour).UnixNano(), 1, &encoder)
	decoder := binary.HeaderDecoder{}

	assert.Nil(t, decoder.Decode(&buf, decodedHeader))

	assert.Equal(t, *header, *decodedHeader)
}
