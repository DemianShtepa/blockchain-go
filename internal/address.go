package internal

import (
	"encoding/hex"
	"fmt"
)

const addressLength = 20

type Address [addressLength]byte

func (a *Address) toSlice() []byte {
	slice := make([]byte, addressLength)

	for i, b := range a {
		slice[i] = b
	}

	return slice
}

func (a *Address) String() string {
	return hex.EncodeToString(a.toSlice())
}

func AddressFromBytes(b []byte) (Address, error) {
	if len(b) != addressLength {
		return Address{}, fmt.Errorf("expected bytes length to be %d, got %d", addressLength, len(b))
	}

	var address Address
	for i := range address {
		address[i] = b[i]
	}

	return address, nil
}
