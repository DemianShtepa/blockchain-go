package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/DemianShtepa/blockchain-go/internal"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func NewPrivateKey() (PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return PrivateKey{}, err
	}

	return PrivateKey{
		key: key,
	}, nil
}

func (pk PrivateKey) Sign(data []byte) (*Signature, error) {
	signatureData, err := ecdsa.SignASN1(rand.Reader, pk.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{
		data: signatureData,
	}, nil
}

func (pk PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &pk.key.PublicKey,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pk PublicKey) toBytes() []byte {
	return elliptic.MarshalCompressed(pk.key, pk.key.X, pk.key.Y)
}

func (pk PublicKey) Address() (internal.Address, error) {
	hash := sha256.New()
	hash.Write(pk.toBytes())

	return internal.AddressFromBytes(hash.Sum(nil)[:20])
}

type Signature struct {
	data []byte
}

func (s Signature) Verify(pk PublicKey, data []byte) bool {
	return ecdsa.VerifyASN1(pk.key, data, s.data)
}
