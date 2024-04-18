package block

import (
	"github.com/DemianShtepa/blockchain-go/internal"
)

type Header struct {
	Version       uint64
	DataHash      internal.Hash
	PreviousBlock internal.Hash
	Timestamp     int64
	Height        uint64
}

func NewHeader(version uint64, previousBlock internal.Hash, timestamp int64, height uint64) *Header {
	return &Header{Version: version, PreviousBlock: previousBlock, Timestamp: timestamp, Height: height}
}
