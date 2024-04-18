package block

import (
	"bytes"
	"github.com/DemianShtepa/blockchain-go/internal"
	"io"
)

type BlockEncoder interface {
	Encode(io.Writer, *Block) error
}

type Block struct {
	Head         *Header
	Transactions Transactions

	headerHash internal.Hash
	publicKey  internal.PublicKey
	signature  internal.Signature
}

func (b *Block) PublicKey() internal.PublicKey {
	return b.publicKey
}

func (b *Block) Signature() internal.Signature {
	return b.signature
}

func NewBlock(head *Header, transactions []Transaction) *Block {
	return &Block{Head: head, Transactions: transactions}
}

func (b *Block) Sign(privateKey internal.PrivateKey) error {
	signature, err := privateKey.Sign(b.headerHash[:])
	if err != nil {
		return err
	}

	b.signature = *signature
	b.publicKey = privateKey.PublicKey()

	return nil
}

func (b *Block) Verify() bool {
	return b.signature.Verify(b.publicKey, b.headerHash[:])
}

func (b *Block) Hash(encoder BlockEncoder) (internal.Hash, error) {
	if b.headerHash.IsEmpty() {
		var buf bytes.Buffer

		if err := encoder.Encode(&buf, b); err != nil {
			return internal.Hash{}, err
		}

		headerHash, err := internal.HashFromReader(&buf)
		if err != nil {
			return internal.Hash{}, err
		}

		b.headerHash = headerHash
	}

	return b.headerHash, nil
}
