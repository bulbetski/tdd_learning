package sync

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.value++
	defer c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}
