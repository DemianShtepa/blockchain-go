package block

import (
	"github.com/DemianShtepa/blockchain-go/internal"
)

type Transactions []Transaction

type Transaction struct {
	Data []byte

	PublicKey internal.PublicKey
	Signature internal.Signature
}

func (t *Transaction) Sign(privateKey internal.PrivateKey) error {
	signature, err := privateKey.Sign(t.Data)
	if err != nil {
		return err
	}

	t.Signature = *signature
	t.PublicKey = privateKey.PublicKey()

	return nil
}

func (t *Transaction) Verify() bool {
	return t.Signature.Verify(t.PublicKey, t.Data)
}
