// rate_limiting_02: Allowing bursts (reference solution)
package limiter

import "time"

// Limiter allows bursts of up to burst events, refilled at one per every.
type Limiter struct {
	tokens chan struct{}
	stop   chan struct{}
}

// NewLimiter returns a full limiter and starts its refill goroutine.
func NewLimiter(burst int, every time.Duration) *Limiter {
	l := &Limiter{tokens: make(chan struct{}, burst), stop: make(chan struct{})}
	for range burst {
		l.tokens <- struct{}{}
	}
	go func() {
		ticker := time.NewTicker(every)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case l.tokens <- struct{}{}:
				default: // already full: drop the token
				}
			case <-l.stop:
				return
			}
		}
	}()
	return l
}

// Wait blocks until a token is available and takes it.
func (l *Limiter) Wait() { <-l.tokens }

// Stop stops refilling.
func (l *Limiter) Stop() { close(l.stop) }
