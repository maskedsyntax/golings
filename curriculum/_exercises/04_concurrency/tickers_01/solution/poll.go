// tickers_01: Polling on a schedule (reference solution)
package poll

import "time"

// Poll calls check once per interval (waiting before each call) until it
// returns true, and returns the number of calls.
func Poll(interval time.Duration, check func() bool) int {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	n := 0
	for range ticker.C {
		n++
		if check() {
			return n
		}
	}
	return n // unreachable: ticker.C is never closed
}
