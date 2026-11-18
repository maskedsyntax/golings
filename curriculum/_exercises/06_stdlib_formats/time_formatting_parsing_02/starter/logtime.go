// time_formatting_parsing_02: Two log formats
//
// Implement ParseLogTime. See README.md.
//
// Reference: https://gobyexample.com/time-formatting-parsing
package logtime

import (
	"errors"
	"time"
)

// ParseLogTime parses an RFC 3339 or access-log timestamp and returns it in UTC.
func ParseLogTime(s string) (time.Time, error) {
	// TODO
	return time.Time{}, errors.New("not implemented")
}
