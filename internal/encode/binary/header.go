package binary

import (
	"encoding/binary"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"io"
)

type HeaderEncoder struct {
}

func (he *HeaderEncoder) Encode(writer io.Writer, header *block.Header) error {
	if err := binary.Write(writer, binary.LittleEndian, header.Version); err != nil {
		return err
	}

	if err := binary.Write(writer, binary.LittleEndian, header.PreviousBlock); err != nil {
		return err
	}

	if err := binary.Write(writer, binary.LittleEndian, header.Timestamp); err != nil {
		return err
	}

	return binary.Write(writer, binary.LittleEndian, header.Height)
}

type HeaderDecoder struct {
}

func (hd *HeaderDecoder) Decode(reader io.Reader, header *block.Header) error {
	if err := binary.Read(reader, binary.LittleEndian, &header.Version); err != nil {
		return err
	}

	if err := binary.Read(reader, binary.LittleEndian, &header.PreviousBlock); err != nil {
		return err
	}

	if err := binary.Read(reader, binary.LittleEndian, &header.Timestamp); err != nil {
		return err
	}

	return binary.Read(reader, binary.LittleEndian, &header.Height)
}
