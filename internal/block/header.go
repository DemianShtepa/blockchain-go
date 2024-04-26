package block

import (
	"bytes"
	"github.com/DemianShtepa/blockchain-go/internal"
	"io"
)

type HeaderEncoder interface {
	Encode(io.Writer, *Header) error
}

type Header struct {
	Version           uint64
	DataHash          internal.Hash
	PreviousBlockHash internal.Hash
	Timestamp         int64
	Height            uint64

	headerEncoder HeaderEncoder
}

func NewHeader(
	version uint64,
	previousBlockHash internal.Hash,
	timestamp int64,
	height uint64,
	headerEncoder HeaderEncoder,
) *Header {
	return &Header{
		Version:           version,
		PreviousBlockHash: previousBlockHash,
		Timestamp:         timestamp,
		Height:            height,
		headerEncoder:     headerEncoder,
	}
}

func (header *Header) Hash() (internal.Hash, error) {
	var buf bytes.Buffer

	if err := header.headerEncoder.Encode(&buf, header); err != nil {
		return internal.Hash{}, err
	}

	return internal.HashFromBytes(buf.Bytes()), nil
}
