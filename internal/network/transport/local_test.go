package transport_test

import (
	"github.com/DemianShtepa/blockchain-go/internal/network/transport"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestLocalTransport_Connect(t *testing.T) {
	localTransportA := transport.NewLocalTransport("A")
	localTransportB := transport.NewLocalTransport("B")

	assert.Nil(t, localTransportA.Connect(localTransportB))
}

func TestLocalTransport_Connect_FailWithInvalidTransport(t *testing.T) {
	localTransportA := transport.NewLocalTransport("A")

	assert.NotNil(t, localTransportA.Connect(nil))
}

func TestLocalTransport_Connect_FailWithTransportDuplication(t *testing.T) {
	localTransportA := transport.NewLocalTransport("A")
	localTransportB := transport.NewLocalTransport("B")

	assert.Nil(t, localTransportA.Connect(localTransportB))
	assert.NotNil(t, localTransportA.Connect(localTransportB))
}

func TestLocalTransport_SendMessage(t *testing.T) {
	localTransportA := transport.NewLocalTransport("A")
	localTransportB := transport.NewLocalTransport("B")

	assert.Nil(t, localTransportA.Connect(localTransportB))

	assert.Nil(t, localTransportA.SendMessage(localTransportB.Address(), []byte("test")))
}

func TestLocalTransport_SendMessage_FailWithNonExistingTransport(t *testing.T) {
	localTransportA := transport.NewLocalTransport("A")
	localTransportB := transport.NewLocalTransport("B")

	assert.NotNil(t, localTransportA.SendMessage(localTransportB.Address(), []byte("test")))
}

func TestLocalTransport_ConcurrentAccess(t *testing.T) {
	localTransportA := transport.NewLocalTransport("A")
	localTransportB := transport.NewLocalTransport("B")

	assert.Nil(t, localTransportA.Connect(localTransportB))

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		localTransportC := transport.NewLocalTransport("C")

		assert.Nil(t, localTransportA.Connect(localTransportC))
	}()

	go func() {
		defer wg.Done()

		assert.Nil(t, localTransportA.SendMessage(localTransportB.Address(), []byte("test")))
	}()

	wg.Wait()
}
