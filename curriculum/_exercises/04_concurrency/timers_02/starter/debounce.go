// timers_02: Debouncing
//
// f runs once per trigger instead of once per burst. Fix it. See README.md.
//
// Reference: https://gobyexample.com/timers
package debounce

import "time"

// Debouncer runs f once, d after the most recent Trigger.
type Debouncer struct {
	d time.Duration
	f func()
}

// NewDebouncer returns a Debouncer for f.
func NewDebouncer(d time.Duration, f func()) *Debouncer {
	return &Debouncer{d: d, f: f}
}

// Trigger records an event. Call it from a single goroutine.
func (db *Debouncer) Trigger() {
	time.AfterFunc(db.d, db.f)
}
