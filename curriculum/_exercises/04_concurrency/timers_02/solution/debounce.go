// timers_02: Debouncing (reference solution)
package debounce

import "time"

// Debouncer runs f once, d after the most recent Trigger.
type Debouncer struct {
	d time.Duration
	t *time.Timer
}

// NewDebouncer returns a Debouncer for f.
func NewDebouncer(d time.Duration, f func()) *Debouncer {
	t := time.AfterFunc(d, f)
	t.Stop() // nothing scheduled until the first Trigger
	return &Debouncer{d: d, t: t}
}

// Trigger records an event. Call it from a single goroutine.
func (db *Debouncer) Trigger() {
	db.t.Reset(db.d) // push the pending call back to d from now
}
