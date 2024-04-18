package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
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

	copy(slice, h[:HashLength])

	return slice
}

func (h *Hash) String() string {
	return hex.EncodeToString(h.toSlice())
}

func HashFromReader(reader io.Reader) (Hash, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(reader)
	if err != nil {
		return Hash{}, err
	}

	return HashFromBytes(buf.Bytes())
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
