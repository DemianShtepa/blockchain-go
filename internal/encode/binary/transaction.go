package binary

import (
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"io"
)

type TransactionEncoder struct {
}

func (te *TransactionEncoder) Encode(writer io.Writer, transactions ...block.Transaction) error {
	return nil
}

type TransactionDecoder struct {
}

func (td *TransactionDecoder) Decode(reader io.Reader, transaction block.Transaction) error {
	return nil
}
