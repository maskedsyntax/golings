// mutexes_04: Compute once per key (reference solution)
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
		return v // fast path: a read lock only
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.data[key]; ok {
		return v // someone else computed it while we waited for the lock
	}
	v = compute()
	c.data[key] = v
	return v
}
