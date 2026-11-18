// epoch_01: Milliseconds since 1970 (reference solution)
package epoch

import "time"

// ToUnixMillis returns t as milliseconds since the Unix epoch.
func ToUnixMillis(t time.Time) int64 {
	return t.UnixMilli()
}

// FromUnixMillis converts milliseconds since the Unix epoch to a UTC time.
func FromUnixMillis(ms int64) time.Time {
	return time.UnixMilli(ms).UTC()
}
