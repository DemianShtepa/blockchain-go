package encode

import (
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"io"
)

type BlockEncoder struct {
	writer             io.Writer
	headerEncoder      *HeaderEncoder
	transactionEncoder *TransactionEncoder
}

func NewBlockEncoder(
	writer io.Writer,
	headerEncoder *HeaderEncoder,
	transactionEncoder *TransactionEncoder,
) *BlockEncoder {
	return &BlockEncoder{
		writer:             writer,
		headerEncoder:      headerEncoder,
		transactionEncoder: transactionEncoder,
	}
}

func (be *BlockEncoder) Encode(b *block.Block) error {
	if err := be.headerEncoder.Encode(b.Head); err != nil {
		return err
	}

	return be.transactionEncoder.Encode(b.Transactions...)
}

type BlockDecoder struct {
	reader             io.Reader
	headerDecoder      *HeaderDecoder
	transactionDecoder *TransactionDecoder
}

func NewBlockDecoder(
	reader io.Reader,
	headerDecoder *HeaderDecoder,
	transactionDecoder *TransactionDecoder,
) *BlockDecoder {
	return &BlockDecoder{
		reader:             reader,
		headerDecoder:      headerDecoder,
		transactionDecoder: transactionDecoder,
	}
}

func (bd *BlockDecoder) Decode(b *block.Block) error {
	if err := bd.headerDecoder.Decode(b.Head); err != nil {
		return err
	}

	return nil
}
