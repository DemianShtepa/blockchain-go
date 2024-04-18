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

	hash, _ := randomHash(t)
	header := block.NewHeader(1, hash, time.Now().UnixNano(), 10)
	encoder := binary.HeaderEncoder{}

	assert.Nil(t, encoder.Encode(&buf, header))

	var decodedHeader block.Header
	decoder := binary.HeaderDecoder{}

	assert.Nil(t, decoder.Decode(&buf, &decodedHeader))

	assert.Equal(t, *header, decodedHeader)
}
