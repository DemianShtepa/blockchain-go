package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressFromBytes(t *testing.T) {
	cases := []struct {
		title       string
		bytesLength int
		assertion   func(t *testing.T, err error)
	}{
		{
			title:       "wrong address length",
			bytesLength: 1,
			assertion: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			title:       "correct address length",
			bytesLength: addressLength,
			assertion: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			b := make([]byte, c.bytesLength)

			_, err := AddressFromBytes(b)

			c.assertion(t, err)
		})
	}
}
