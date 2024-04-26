package network

import (
	"context"
	"errors"
	"fmt"
	"github.com/DemianShtepa/blockchain-go/internal"
	"github.com/DemianShtepa/blockchain-go/internal/block"
	"log/slog"
	"time"
)

type Address string

type RPC struct {
	From    Address
	Payload []byte
}

type Transport interface {
	Connect(Transport) error
	Consume() <-chan RPC
	SendMessage(Address, []byte) error
	Address() Address
}

type ServerOpts struct {
	Transports      []Transport
	TransactionPool *Pool
	PrivateKey      *internal.PrivateKey
	Logger          *slog.Logger
}

type Server struct {
	Transports []Transport

	logger                   *slog.Logger
	transactionFlushDuration time.Duration
	transactionPool          *Pool
	privateKey               *internal.PrivateKey
	rpcChannel               chan RPC
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		Transports:               opts.Transports,
		logger:                   opts.Logger,
		transactionFlushDuration: time.Minute,
		transactionPool:          opts.TransactionPool,
		privateKey:               opts.PrivateKey,
		rpcChannel:               make(chan RPC, 1024),
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.initTransports()
	ticker := time.NewTicker(s.transactionFlushDuration)
	defer ticker.Stop()

	for {
		select {
		case rpc := <-s.rpcChannel:
			fmt.Println(rpc)
		case <-ticker.C:
			if s.isValidator() {
				s.createBlock()
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *Server) handleTransaction(transaction *block.Transaction) error {
	if !transaction.Verify() {
		return errors.New("invalid signature of a transaction")
	}

	err := s.transactionPool.Add(transaction)
	if err != nil {
		if errors.Is(err, ErrTransactionAlreadyExists) {
			hash := transaction.Hash()
			s.logger.Info(err.Error(), slog.String("hash", hash.String()))
		}
	}

	return nil
}

func (s *Server) createBlock() {
}

func (s *Server) isValidator() bool {
	return s.privateKey != nil
}

func (s *Server) initTransports() {
	for _, transport := range s.Transports {
		transport := transport

		go func() {
			for rpc := range transport.Consume() {
				s.rpcChannel <- rpc
			}
		}()
	}
}
