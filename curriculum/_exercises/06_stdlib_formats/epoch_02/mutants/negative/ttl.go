// Mutant: forgets to clamp expired tokens to 0.
package ttl

import "time"

func TTL(expiresAt int64, now time.Time) time.Duration {
	return time.Unix(expiresAt, 0).Sub(now)
}
