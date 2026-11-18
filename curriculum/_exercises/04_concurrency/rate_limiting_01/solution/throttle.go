// rate_limiting_01: One request per tick (reference solution)
package throttle

import "time"

// Throttle calls handle for each request in order, at most once per every.
func Throttle(reqs []string, every time.Duration, handle func(string)) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()
	for _, r := range reqs {
		<-ticker.C
		handle(r)
	}
}
