// timeouts_01: Giving up
//
// Implement WithTimeout. See README.md.
//
// Reference: https://gobyexample.com/timeouts
package timeout

import (
	"errors"
	"time"
)

// ErrTimeout is returned when fetch takes longer than allowed.
var ErrTimeout = errors.New("timed out")

// WithTimeout returns fetch's result, or ErrTimeout if it takes longer than d.
func WithTimeout(fetch func() string, d time.Duration) (string, error) {
	// TODO: run fetch in a goroutine and select on its result and time.After(d).
	return fetch(), nil
}
