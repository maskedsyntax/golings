// Mutant: right instant, but in the local zone.
package epoch

import "time"

func ToUnixMillis(t time.Time) int64 {
	return t.UnixMilli()
}

func FromUnixMillis(ms int64) time.Time {
	return time.UnixMilli(ms)
}
