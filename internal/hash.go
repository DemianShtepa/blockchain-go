package internal

import (
	"encoding/hex"
	"fmt"
)

const hashLength = 32

type Hash [hashLength]byte

func (h *Hash) IsEmpty() bool {
	for _, u := range h {
		if u != 0 {
			return false
		}
	}

	return true
}

func (h *Hash) toSlice() []byte {
	slice := make([]byte, hashLength)
	for i, u := range h {
		slice[i] = u
	}

	return slice
}

func (h *Hash) String() string {
	return hex.EncodeToString(h.toSlice())
}

func HashFromBytes(b []byte) (Hash, error) {
	if len(b) != hashLength {
		return Hash{}, fmt.Errorf("expected bytes length to be %d, got %d", hashLength, len(b))
	}

	var hash Hash
	for i, bt := range b {
		hash[i] = bt
	}

	return hash, nil
}
