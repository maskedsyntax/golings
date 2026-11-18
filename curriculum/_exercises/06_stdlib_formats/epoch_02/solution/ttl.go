// epoch_02: Time to live (reference solution)
package ttl

import "time"

// TTL returns how long remains until expiresAt (Unix seconds), or 0 if it has passed.
func TTL(expiresAt int64, now time.Time) time.Duration {
	left := time.Unix(expiresAt, 0).Sub(now)
	if left < 0 {
		return 0
	}
	return left
}
