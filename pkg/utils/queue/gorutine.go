package queue

import (
	"fmt"
	"sync"
	"time"
)

type GoroutineDriver struct {
	queues map[string]chan string
	mu     sync.RWMutex
}

func NewGoroutineDriver(bufferSize int) *GoroutineDriver {
	return &GoroutineDriver{
		queues: make(map[string]chan string),
	}
}

func (g *GoroutineDriver) getQueue(name string) chan string {
	g.mu.Lock()
	defer g.mu.Unlock()
	if q, ok := g.queues[name]; ok {
		return q
	}
	q := make(chan string, 100)
	g.queues[name] = q
	return q
}

func (g *GoroutineDriver) Push(name string, payload any) error {
	q := g.getQueue(name)
	payloadStr := fmt.Sprintf("%v", payload)
	q <- payloadStr
	return nil
}

func (g *GoroutineDriver) Later(name string, payload any, delay time.Duration) error {
	go func() {
		time.Sleep(delay)
		g.Push(name, payload)
	}()
	return nil
}

func (g *GoroutineDriver) Pop(name string) (string, error) {
	q := g.getQueue(name)
	select {
	case payload := <-q:
		return payload, nil
	default:
		return "", nil
	}
}
