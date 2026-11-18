// mutexes_01: Guarding a map (reference solution)
package container

import "sync"

// Container holds named counters. It is safe for concurrent use.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// New returns an empty Container.
func New() *Container { return &Container{counters: map[string]int{}} }

// Inc adds one to the named counter.
func (c *Container) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// Get returns the named counter.
func (c *Container) Get(name string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counters[name]
}
