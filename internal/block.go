package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"
)

type Header struct {
	Version       uint64
	PreviousBlock Hash
	Timestamp     int64
	Height        uint64
	Nonce         uint64
}

func NewHeader(version uint64, previousBlock Hash, timestamp int64, height uint64, nonce uint64) *Header {
	return &Header{Version: version, PreviousBlock: previousBlock, Timestamp: timestamp, Height: height, Nonce: nonce}
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, h.Version); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, h.PreviousBlock); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, h.Timestamp); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, h.Height); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, h.Nonce)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &h.PreviousBlock); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

type Block struct {
	Head         *Header
	Transactions Transactions

	headerHash Hash
}

func NewBlock(head *Header, transactions []Transaction) *Block {
	return &Block{Head: head, Transactions: transactions}
}

func (b *Block) Hash() (Hash, error) {
	buf := bytes.Buffer{}
	if err := b.EncodeBinary(&buf); err != nil {
		return Hash{}, err
	}

	shaHash := sha256.New()
	shaHash.Write(buf.Bytes())
	hash, err := HashFromBytes(shaHash.Sum(nil))
	if err != nil {
		return Hash{}, err
	}

	return hash, nil
}

func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Head.EncodeBinary(w); err != nil {
		return err
	}

	return b.Transactions.EncodeBinary(w)
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Head.DecodeBinary(r); err != nil {
		return err
	}

	return b.Transactions.DecodeBinary(r)
}
