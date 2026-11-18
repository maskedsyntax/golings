// Mutant: cancel always claims success.
package schedule

import "time"

func Schedule(d time.Duration, f func()) (cancel func() bool) {
	t := time.AfterFunc(d, f)
	return func() bool { t.Stop(); return true }
}
