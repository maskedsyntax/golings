// timeouts_01: Giving up (reference solution)
package timeout

import (
	"errors"
	"time"
)

// ErrTimeout is returned when fetch takes longer than allowed.
var ErrTimeout = errors.New("timed out")

// WithTimeout returns fetch's result, or ErrTimeout if it takes longer than d.
func WithTimeout(fetch func() string, d time.Duration) (string, error) {
	ch := make(chan string, 1) // buffered: a late fetch can still send and exit
	go func() { ch <- fetch() }()
	select {
	case v := <-ch:
		return v, nil
	case <-time.After(d):
		return "", ErrTimeout
	}
}
