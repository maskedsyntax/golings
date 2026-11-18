// epoch_01: Milliseconds since 1970
//
// The conversions lose milliseconds. Fix them. See README.md.
//
// Reference: https://gobyexample.com/epoch
package epoch

import "time"

// ToUnixMillis returns t as milliseconds since the Unix epoch.
func ToUnixMillis(t time.Time) int64 {
	return t.Unix() * 1000
}

// FromUnixMillis converts milliseconds since the Unix epoch to a UTC time.
func FromUnixMillis(ms int64) time.Time {
	return time.Unix(ms/1000, 0)
}
