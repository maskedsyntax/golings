// mutexes_02: The copied lock
//
// The count never changes. Fix Counter. See README.md.
//
// Reference: https://gobyexample.com/mutexes
package counter

import "sync"

// Counter is a mutex-protected counter.
type Counter struct {
	mu sync.Mutex
	n  int
}

// Inc adds one.
func (c Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

// Value returns the count.
func (c Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}
