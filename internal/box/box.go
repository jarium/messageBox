package box

import (
	"github.com/jarium/messageBox/pkg/connector"
	"sync"
)

type Box struct {
	mu       sync.RWMutex
	cond     *sync.Cond
	messages []*connector.Message
}

func New() *Box {
	b := &Box{
		messages: make([]*connector.Message, 0),
	}

	b.cond = sync.NewCond(&b.mu)

	return b
}

func (b *Box) Push(m *connector.Message) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.messages = append(b.messages, m)
	b.cond.Signal()
}

func (b *Box) Pop() *connector.Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.messages) == 0 {
		b.cond.Wait()
	}

	m := b.messages[0]
	b.messages = b.messages[1:]

	return m
}
