// time_formatting_parsing_02: Two log formats (reference solution)
package logtime

import (
	"fmt"
	"time"
)

const accessLog = "02/Jan/2006:15:04:05 -0700"

// ParseLogTime parses an RFC 3339 or access-log timestamp and returns it in UTC.
func ParseLogTime(s string) (time.Time, error) {
	for _, layout := range []string{time.RFC3339, accessLog} {
		if t, err := time.Parse(layout, s); err == nil {
			return t.UTC(), nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized timestamp %q", s)
}
