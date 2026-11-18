// tickers_02: A stoppable heartbeat (reference solution)
package heartbeat

import "time"

// Heartbeat calls beat every interval until stop is closed.
func Heartbeat(interval time.Duration, stop <-chan struct{}, beat func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			beat()
		case <-stop:
			return
		}
	}
}
