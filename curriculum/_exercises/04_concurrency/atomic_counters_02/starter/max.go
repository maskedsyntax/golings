// atomic_counters_02: Compare and swap
//
// Implement Observe and Value. See README.md.
//
// Reference: https://gobyexample.com/atomic-counters
package max

import "sync/atomic"

// Max tracks the largest value observed. It is safe for concurrent use.
type Max struct {
	v atomic.Int64
}

// Observe records x.
func (m *Max) Observe(x int64) {
	// TODO
}

// Value returns the largest value observed so far (0 if none).
func (m *Max) Value() int64 {
	// TODO
	return -1
}
