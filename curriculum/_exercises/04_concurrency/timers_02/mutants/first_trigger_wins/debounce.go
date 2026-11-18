// Mutant: schedules on the first trigger of a burst and ignores the rest.
package debounce

import (
	"sync/atomic"
	"time"
)

type Debouncer struct {
	d       time.Duration
	f       func()
	pending atomic.Bool
}

func NewDebouncer(d time.Duration, f func()) *Debouncer {
	return &Debouncer{d: d, f: f}
}

func (db *Debouncer) Trigger() {
	if db.pending.CompareAndSwap(false, true) {
		time.AfterFunc(db.d, func() { db.pending.Store(false); db.f() })
	}
}
