package internal

import (
	"io"
)

type Transactions []Transaction

func (t Transactions) EncodeBinary(w io.Writer) error {
	for _, transaction := range t {
		if err := transaction.EncodeBinary(w); err != nil {
			return err
		}
	}

	return nil
}

func (t Transactions) DecodeBinary(r io.Reader) error {
	for _, transaction := range t {
		if err := transaction.DecodeBinary(r); err != nil {
			return err
		}
	}

	return nil
}

type Transaction struct {
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}
