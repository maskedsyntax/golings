// Mutant: only understands RFC 3339.
package logtime

import "time"

func ParseLogTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	return t.UTC(), err
}
