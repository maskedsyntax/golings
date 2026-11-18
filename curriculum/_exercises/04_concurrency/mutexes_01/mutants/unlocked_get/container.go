// Mutant: Get reads without the lock.
package container

import "sync"

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func New() *Container { return &Container{counters: map[string]int{}} }

func (c *Container) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func (c *Container) Get(name string) int {
	return c.counters[name]
}
