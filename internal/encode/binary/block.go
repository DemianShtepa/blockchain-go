package binary

import (
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"io"
)

type BlockEncoder struct {
	headerEncoder      HeaderEncoder
	transactionEncoder TransactionEncoder
}

func (be *BlockEncoder) Encode(w io.Writer, b *block.Block) error {
	if err := be.headerEncoder.Encode(w, b.Head); err != nil {
		return err
	}

	return be.transactionEncoder.Encode(w, b.Transactions...)
}

type BlockDecoder struct {
	headerDecoder HeaderDecoder
}

func (bd *BlockDecoder) Decode(r io.Reader, b *block.Block) error {
	if err := bd.headerDecoder.Decode(r, b.Head); err != nil {
		return err
	}

	return nil
}
