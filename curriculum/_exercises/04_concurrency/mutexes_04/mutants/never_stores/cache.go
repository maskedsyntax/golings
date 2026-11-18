// Mutant: holds the lock while computing but never stores the value.
package cache

import "sync"

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Cache { return &Cache{data: map[string]string{}} }

func (c *Cache) Get(key string, compute func() string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.data[key]; ok {
		return v
	}
	return compute()
}
