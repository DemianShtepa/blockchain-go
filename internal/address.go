package internal

import (
	"encoding/hex"
	"fmt"
)

const AddressLength = 20

type Address [AddressLength]byte

func (a *Address) toSlice() []byte {
	slice := make([]byte, AddressLength)

	copy(slice, a[:AddressLength])

	return slice
}

func (a *Address) String() string {
	return hex.EncodeToString(a.toSlice())
}

func AddressFromBytes(b []byte) (Address, error) {
	if len(b) != AddressLength {
		return Address{}, fmt.Errorf("expected bytes length to be %d, got %d", AddressLength, len(b))
	}

	var address Address

	copy(address[:], b[:AddressLength])

	return address, nil
}
