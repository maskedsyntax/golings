// Mutant: waits once, then handles the rest immediately.
package throttle

import "time"

func Throttle(reqs []string, every time.Duration, handle func(string)) {
	time.Sleep(every)
	for _, r := range reqs {
		handle(r)
	}
}
