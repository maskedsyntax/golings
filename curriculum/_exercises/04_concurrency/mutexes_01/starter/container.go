// mutexes_01: Guarding a map
//
// Container isn't safe for concurrent use. Use its mutex. See README.md.
//
// Reference: https://gobyexample.com/mutexes
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
	c.counters[name]++
}

// Get returns the named counter.
func (c *Container) Get(name string) int {
	return c.counters[name]
}
