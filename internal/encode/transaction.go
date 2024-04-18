package encode

import (
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"io"
)

type TransactionEncoder struct {
	writer io.Writer
}

func NewTransactionEncoder(writer io.Writer) *TransactionEncoder {
	return &TransactionEncoder{writer: writer}
}

func (te *TransactionEncoder) Encode(transactions ...block.Transaction) error {
	return nil
}

type TransactionDecoder struct {
	reader io.Reader
}

func NewTransactionDecoder(reader io.Reader) *TransactionDecoder {
	return &TransactionDecoder{reader: reader}
}

func (td *TransactionDecoder) Decode(transaction block.Transaction) error {
	return nil
}
