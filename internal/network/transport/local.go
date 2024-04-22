package transport

import (
	"errors"
	"github.com/DemianShtepa/blockchain-go/internal/network"
	"sync"
)

type LocalTransport struct {
	address        network.Address
	consumeChannel chan network.RPC
	lock           sync.RWMutex
	peers          map[network.Address]*LocalTransport
}

func NewLocalTransport(address network.Address) *LocalTransport {
	return &LocalTransport{
		address:        address,
		consumeChannel: make(chan network.RPC, 1024),
		peers:          make(map[network.Address]*LocalTransport),
	}
}

func (l *LocalTransport) Connect(transport network.Transport) error {
	localTransport, ok := transport.(*LocalTransport)
	if !ok {
		return errors.New("invalid type of transport provided")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	address := localTransport.Address()
	_, ok = l.peers[address]
	if ok {
		return errors.New("transport already connected")
	}

	l.peers[address] = localTransport

	return nil
}

func (l *LocalTransport) Consume() <-chan network.RPC {
	return l.consumeChannel
}

func (l *LocalTransport) SendMessage(to network.Address, payload []byte) error {
	l.lock.RLock()
	defer l.lock.RUnlock()

	transport, ok := l.peers[to]
	if !ok {
		return errors.New("transport doesn't exist")
	}

	transport.consumeChannel <- network.RPC{
		From:    l.Address(),
		Payload: payload,
	}

	return nil
}

func (l *LocalTransport) Address() network.Address {
	return l.address
}
