// timers_01: A cancellable reminder (reference solution)
package schedule

import "time"

// Schedule runs f after d unless the returned cancel is called first.
// cancel reports whether it stopped f from running.
func Schedule(d time.Duration, f func()) (cancel func() bool) {
	t := time.AfterFunc(d, f)
	return t.Stop // a method value: calling it calls t.Stop()
}
