// mutexes_04: Compute once per key
//
// Concurrent misses compute the same key many times. Fix Get. See README.md.
//
// Reference: https://gobyexample.com/mutexes
package cache

import "sync"

// Cache memoizes values by key. It is safe for concurrent use.
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

// New returns an empty Cache.
func New() *Cache { return &Cache{data: map[string]string{}} }

// Get returns the value for key, calling compute (once) if it is missing.
func (c *Cache) Get(key string, compute func() string) string {
	c.mu.RLock()
	v, ok := c.data[key]
	c.mu.RUnlock()
	if ok {
		return v
	}
	v = compute()
	c.mu.Lock()
	c.data[key] = v
	c.mu.Unlock()
	return v
}
