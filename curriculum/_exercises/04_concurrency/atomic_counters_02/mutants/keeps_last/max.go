// Mutant: stores the most recent value instead of the largest.
package max

import "sync/atomic"

type Max struct {
	v atomic.Int64
}

func (m *Max) Observe(x int64) { m.v.Store(x) }

func (m *Max) Value() int64 { return m.v.Load() }
