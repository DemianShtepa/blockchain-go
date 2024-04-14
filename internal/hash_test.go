package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash_IsEmpty(t *testing.T) {
	cases := []struct {
		title   string
		hash    Hash
		isEmpty bool
	}{
		{
			title:   "empty hash",
			hash:    Hash{},
			isEmpty: true,
		},
		{
			title:   "not empty hash",
			hash:    Hash{1},
			isEmpty: false,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			assert.Equal(t, c.isEmpty, c.hash.IsEmpty())
		})
	}
}

func TestHashFromBytes_ReturnsError(t *testing.T) {
	b := make([]byte, 1)

	hash, err := HashFromBytes(b)
	assert.True(t, hash.IsEmpty())
	assert.NotNil(t, err)
}

func TestHashFromBytes(t *testing.T) {
	b := make([]byte, hashLength)

	hash, err := HashFromBytes(b)
	assert.True(t, hash.IsEmpty())
	assert.Nil(t, err)
}
