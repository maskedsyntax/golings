// Mutant: reports one call too many.
package poll

import "time"

func Poll(interval time.Duration, check func() bool) int {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	n := 0
	for range ticker.C {
		n++
		if check() {
			return n + 1
		}
	}
	return n
}
