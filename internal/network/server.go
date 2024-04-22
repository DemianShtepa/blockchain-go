package network

import (
	"context"
	"fmt"
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

type Server struct {
	Transports []Transport

	rpcChannel chan RPC
}

func NewServer(transports ...Transport) *Server {
	return &Server{
		Transports: transports,
		rpcChannel: make(chan RPC, 1024),
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.initTransports()

	for {
		select {
		case rpc := <-s.rpcChannel:
			fmt.Println(rpc)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
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
