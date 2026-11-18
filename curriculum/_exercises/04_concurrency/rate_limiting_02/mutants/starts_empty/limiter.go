// Mutant: no initial burst.
package limiter

import "time"

type Limiter struct {
	tokens chan struct{}
	stop   chan struct{}
}

func NewLimiter(burst int, every time.Duration) *Limiter {
	l := &Limiter{tokens: make(chan struct{}, burst), stop: make(chan struct{})}
	go func() {
		ticker := time.NewTicker(every)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case l.tokens <- struct{}{}:
				default:
				}
			case <-l.stop:
				return
			}
		}
	}()
	return l
}

func (l *Limiter) Wait() { <-l.tokens }

func (l *Limiter) Stop() { close(l.stop) }
