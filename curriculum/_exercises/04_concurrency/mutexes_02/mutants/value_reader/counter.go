// Mutant: Inc fixed, but Value still copies the lock (fails vet).
package counter

import "sync"

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c Counter) Value() int {
	return c.n
}
