// Mutant: waits for stop without ever beating.
package heartbeat

import "time"

func Heartbeat(interval time.Duration, stop <-chan struct{}, beat func()) {
	<-stop
}
