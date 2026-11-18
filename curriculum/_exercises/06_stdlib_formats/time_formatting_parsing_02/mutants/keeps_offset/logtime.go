// Mutant: right instant, but keeps the input's offset.
package logtime

import (
	"fmt"
	"time"
)

func ParseLogTime(s string) (time.Time, error) {
	for _, layout := range []string{time.RFC3339, "02/Jan/2006:15:04:05 -0700"} {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized timestamp %q", s)
}
