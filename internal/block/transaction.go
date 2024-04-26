package block

import (
	"github.com/DemianShtepa/blockchain-go/internal"
)

type Transactions []Transaction

type Transaction struct {
	Data []byte

	hash      internal.Hash
	publicKey internal.PublicKey
	signature *internal.Signature
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{Data: data}
}

func (t *Transaction) PublicKey() internal.PublicKey {
	return t.publicKey
}

func (t *Transaction) Signature() *internal.Signature {
	return t.signature
}

func (t *Transaction) Sign(privateKey internal.PrivateKey) error {
	signature, err := privateKey.Sign(t.Data)
	if err != nil {
		return err
	}

	t.signature = signature
	t.publicKey = privateKey.PublicKey()

	return nil
}

func (t *Transaction) Verify() bool {
	if t.signature == nil {
		return false
	}

	return t.signature.Verify(t.publicKey, t.Data)
}

func (t *Transaction) Hash() internal.Hash {
	if t.hash.IsEmpty() {
		t.hash = internal.HashFromBytes(t.Data)
	}

	return t.hash
}
