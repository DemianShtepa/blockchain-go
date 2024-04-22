package block

import (
	"github.com/DemianShtepa/blockchain-go/internal"
)

type Transactions []Transaction

type Transaction struct {
	Data []byte

	publicKey internal.PublicKey
	signature *internal.Signature
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
