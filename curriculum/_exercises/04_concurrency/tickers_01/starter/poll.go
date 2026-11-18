// tickers_01: Polling on a schedule
//
// Poll calls check in a tight loop. Use a ticker. See README.md.
//
// Reference: https://gobyexample.com/tickers
package poll

import "time"

// Poll calls check once per interval (waiting before each call) until it
// returns true, and returns the number of calls.
func Poll(interval time.Duration, check func() bool) int {
	n := 0
	for {
		n++
		if check() {
			return n
		}
	}
}
