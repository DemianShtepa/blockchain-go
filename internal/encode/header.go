package encode

import (
	"encoding/binary"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"io"
)

type HeaderEncoder struct {
	writer io.Writer
}

func NewHeaderEncoder(writer io.Writer) *HeaderEncoder {
	return &HeaderEncoder{writer: writer}
}

func (he *HeaderEncoder) Encode(header *block.Header) error {
	if err := binary.Write(he.writer, binary.LittleEndian, header.Version); err != nil {
		return err
	}

	if err := binary.Write(he.writer, binary.LittleEndian, header.PreviousBlock); err != nil {
		return err
	}

	if err := binary.Write(he.writer, binary.LittleEndian, header.Timestamp); err != nil {
		return err
	}

	return binary.Write(he.writer, binary.LittleEndian, header.Height)
}

type HeaderDecoder struct {
	reader io.Reader
}

func NewHeaderDecoder(reader io.Reader) *HeaderDecoder {
	return &HeaderDecoder{reader: reader}
}

func (hd *HeaderDecoder) Decode(header *block.Header) error {
	if err := binary.Read(hd.reader, binary.LittleEndian, &header.Version); err != nil {
		return err
	}

	if err := binary.Read(hd.reader, binary.LittleEndian, &header.PreviousBlock); err != nil {
		return err
	}

	if err := binary.Read(hd.reader, binary.LittleEndian, &header.Timestamp); err != nil {
		return err
	}

	return binary.Read(hd.reader, binary.LittleEndian, &header.Height)
}
