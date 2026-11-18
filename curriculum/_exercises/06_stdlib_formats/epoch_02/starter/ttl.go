// epoch_02: Time to live
//
// TTL returns nanoseconds where it means seconds. Fix it. See README.md.
//
// Reference: https://gobyexample.com/epoch
package ttl

import "time"

// TTL returns how long remains until expiresAt (Unix seconds), or 0 if it has passed.
func TTL(expiresAt int64, now time.Time) time.Duration {
	return time.Duration(expiresAt - now.Unix())
}
