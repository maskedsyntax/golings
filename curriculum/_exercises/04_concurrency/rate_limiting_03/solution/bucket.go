// rate_limiting_03: A token bucket on paper (reference solution)
package bucket

import "time"

// Bucket is a token bucket holding at most capacity tokens and gaining rate
// tokens per second.
type Bucket struct {
	capacity float64
	rate     float64
	tokens   float64
	last     time.Time
}

// NewBucket returns a full bucket.
func NewBucket(capacity int, ratePerSecond float64, now time.Time) *Bucket {
	return &Bucket{capacity: float64(capacity), rate: ratePerSecond, tokens: float64(capacity), last: now}
}

// Allow reports whether a request at time now may proceed, spending a token if so.
func (b *Bucket) Allow(now time.Time) bool {
	b.tokens = min(b.capacity, b.tokens+now.Sub(b.last).Seconds()*b.rate)
	b.last = now
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}
