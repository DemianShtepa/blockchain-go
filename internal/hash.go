package internal

import (
	"crypto/sha256"
	"encoding/hex"
)

const HashLength = 32

type Hash [HashLength]byte

func (h *Hash) IsEmpty() bool {
	for _, u := range h {
		if u != 0 {
			return false
		}
	}

	return true
}

func (h *Hash) toSlice() []byte {
	slice := make([]byte, HashLength)
	for i, u := range h {
		slice[i] = u
	}

	return slice
}

func (h *Hash) String() string {
	return hex.EncodeToString(h.toSlice())
}

func HashFromBytes(b []byte) (Hash, error) {
	shaHash := sha256.New()
	shaHash.Write(b)
	shaHashBytes := shaHash.Sum(nil)

	var hash Hash
	for i := range hash {
		hash[i] = shaHashBytes[i]
	}

	return hash, nil
}
