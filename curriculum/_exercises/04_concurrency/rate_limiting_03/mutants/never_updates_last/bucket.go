// Mutant: measures elapsed time from creation every time (over-refills).
package bucket

import "time"

type Bucket struct {
	capacity float64
	rate     float64
	tokens   float64
	last     time.Time
}

func NewBucket(capacity int, ratePerSecond float64, now time.Time) *Bucket {
	return &Bucket{capacity: float64(capacity), rate: ratePerSecond, tokens: float64(capacity), last: now}
}

func (b *Bucket) Allow(now time.Time) bool {
	b.tokens = min(b.capacity, b.tokens+now.Sub(b.last).Seconds()*b.rate)
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}
