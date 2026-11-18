// timers_01: A cancellable reminder
//
// Schedule can't be cancelled. Use a timer. See README.md.
//
// Reference: https://gobyexample.com/timers
package schedule

import "time"

// Schedule runs f after d unless the returned cancel is called first.
// cancel reports whether it stopped f from running.
func Schedule(d time.Duration, f func()) (cancel func() bool) {
	go func() {
		time.Sleep(d)
		f()
	}()
	return func() bool { return false }
}
