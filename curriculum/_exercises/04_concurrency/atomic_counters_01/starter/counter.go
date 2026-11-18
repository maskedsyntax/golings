// atomic_counters_01: Counting without losing
//
// Counter loses updates under concurrency. Use sync/atomic. See README.md.
//
// Reference: https://gobyexample.com/atomic-counters
package counter

// Counter is a counter that is safe for concurrent use.
type Counter struct {
	n int64
}

// Inc adds one.
func (c *Counter) Inc() { c.n++ }

// Add adds delta (which may be negative).
func (c *Counter) Add(delta int64) { c.n += delta }

// Value returns the current count.
func (c *Counter) Value() int64 { return c.n }
