package block

import (
	"errors"
	"github.com/DemianShtepa/blockchain-go/internal"
)

type Block struct {
	Head         *Header
	Transactions Transactions

	headerHash internal.Hash
	publicKey  internal.PublicKey
	signature  *internal.Signature
}

func NewBlock(head *Header, transactions []Transaction) *Block {
	return &Block{Head: head, Transactions: transactions}
}

func (b *Block) PublicKey() internal.PublicKey {
	return b.publicKey
}

func (b *Block) Signature() *internal.Signature {
	return b.signature
}

func (b *Block) Sign(privateKey internal.PrivateKey) error {
	if b.headerHash.IsEmpty() {
		return errors.New("can't sign block if header hash is empty")
	}

	signature, err := privateKey.Sign(b.headerHash[:])
	if err != nil {
		return err
	}

	b.signature = signature
	b.publicKey = privateKey.PublicKey()

	return nil
}

func (b *Block) Verify() bool {
	if b.signature == nil {
		return false
	}

	if !b.signature.Verify(b.publicKey, b.headerHash[:]) {
		return false
	}

	for _, transaction := range b.Transactions {
		if !transaction.Verify() {
			return false
		}
	}

	return true
}

func (b *Block) Hash() (internal.Hash, error) {
	if b.headerHash.IsEmpty() {
		headerHash, err := b.Head.Hash()
		if err != nil {
			return headerHash, err
		}

		b.headerHash = headerHash
	}

	return b.headerHash, nil
}
