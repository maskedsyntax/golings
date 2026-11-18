// tickers_02: A stoppable heartbeat
//
// Heartbeat never returns. Fix it. See README.md.
//
// Reference: https://gobyexample.com/tickers
package heartbeat

import "time"

// Heartbeat calls beat every interval until stop is closed.
func Heartbeat(interval time.Duration, stop <-chan struct{}, beat func()) {
	for range time.Tick(interval) {
		beat()
	}
}
