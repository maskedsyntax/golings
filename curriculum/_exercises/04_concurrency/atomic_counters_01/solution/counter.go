// atomic_counters_01: Counting without losing (reference solution)
package counter

import "sync/atomic"

// Counter is a counter that is safe for concurrent use.
type Counter struct {
	n atomic.Int64
}

// Inc adds one.
func (c *Counter) Inc() { c.n.Add(1) }

// Add adds delta (which may be negative).
func (c *Counter) Add(delta int64) { c.n.Add(delta) }

// Value returns the current count.
func (c *Counter) Value() int64 { return c.n.Load() }
