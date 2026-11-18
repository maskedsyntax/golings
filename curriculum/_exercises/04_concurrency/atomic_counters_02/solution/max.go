// atomic_counters_02: Compare and swap (reference solution)
package max

import "sync/atomic"

// Max tracks the largest value observed. It is safe for concurrent use.
type Max struct {
	v atomic.Int64
}

// Observe records x.
func (m *Max) Observe(x int64) {
	for {
		cur := m.v.Load()
		if x <= cur || m.v.CompareAndSwap(cur, x) {
			return
		}
		// Another goroutine changed v between Load and CompareAndSwap; retry.
	}
}

// Value returns the largest value observed so far (0 if none).
func (m *Max) Value() int64 { return m.v.Load() }
