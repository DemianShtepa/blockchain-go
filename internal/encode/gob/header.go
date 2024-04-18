package gob

import (
	"encoding/gob"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"io"
)

type HeaderEncoder struct {
}

func (he *HeaderEncoder) Encode(writer io.Writer, header *block.Header) error {
	encoder := gob.NewEncoder(writer)

	return encoder.Encode(*header)
}

type HeaderDecoder struct {
}

func (hd *HeaderDecoder) Decode(reader io.Reader, header *block.Header) error {
	decoder := gob.NewDecoder(reader)

	return decoder.Decode(header)
}
