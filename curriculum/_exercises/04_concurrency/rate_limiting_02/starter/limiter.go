// rate_limiting_02: Allowing bursts
//
// Implement the Limiter. See README.md.
//
// Reference: https://gobyexample.com/rate-limiting
package limiter

import "time"

// Limiter allows bursts of up to burst events, refilled at one per every.
type Limiter struct {
	tokens chan struct{}
	stop   chan struct{}
}

// NewLimiter returns a full limiter and starts its refill goroutine.
func NewLimiter(burst int, every time.Duration) *Limiter {
	// TODO
	return &Limiter{}
}

// Wait blocks until a token is available and takes it.
func (l *Limiter) Wait() {
	// TODO
}

// Stop stops refilling.
func (l *Limiter) Stop() {
	// TODO
}
