package box

import (
	"sync"
)

type Box struct {
	mu       sync.RWMutex
	messages []Message
}

func New() *Box {
	return &Box{
		messages: make([]Message, 0),
	}
}

func (b *Box) Push(m Message) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.messages = append(b.messages, m)
}

func (b *Box) Pop() Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.messages) == 0 {
		return Message{}
	}

	m := b.messages[0]
	b.messages = b.messages[1:]

	return m
}
