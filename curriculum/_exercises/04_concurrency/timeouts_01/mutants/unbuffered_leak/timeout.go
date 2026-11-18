// Mutant: unbuffered result channel leaks the goroutine after a timeout.
package timeout

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timed out")

func WithTimeout(fetch func() string, d time.Duration) (string, error) {
	ch := make(chan string)
	go func() { ch <- fetch() }()
	select {
	case v := <-ch:
		return v, nil
	case <-time.After(d):
		return "", ErrTimeout
	}
}
