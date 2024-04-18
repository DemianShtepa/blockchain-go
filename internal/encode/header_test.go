package encode

import (
	"bytes"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHeader_EncodeDecode(t *testing.T) {
	var buf bytes.Buffer

	hash, _ := randomHash()
	header := block.NewHeader(1, hash, time.Now().UnixNano(), 10)
	encoder := NewHeaderEncoder(&buf)

	assert.Nil(t, encoder.Encode(header))

	var decodedHeader block.Header
	decoder := NewHeaderDecoder(&buf)

	assert.Nil(t, decoder.Decode(&decodedHeader))

	assert.Equal(t, *header, decodedHeader)
}
