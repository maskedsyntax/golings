package bucket

import (
	"testing"
	"time"
)

var t0 = time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

func at(d time.Duration) time.Time { return t0.Add(d) }

func TestBurstAndRefill(t *testing.T) {
	b := NewBucket(2, 1, t0) // 2 tokens, 1 token per second
	steps := []struct {
		at   time.Duration
		want bool
	}{
		{0, true}, {0, true}, {0, false}, // the initial burst of 2
		{500 * time.Millisecond, false}, // half a token
		{time.Second, true},             // one whole token
		{time.Second, false},
		{3 * time.Second, true}, {3 * time.Second, true}, {3 * time.Second, false},
	}
	for _, s := range steps {
		if got := b.Allow(at(s.at)); got != s.want {
			t.Fatalf("Allow at t+%v = %v, want %v", s.at, got, s.want)
		}
	}
}

func TestLongIdleIsCapped(t *testing.T) {
	b := NewBucket(3, 1, t0)
	allowed := 0
	for range 100 {
		if b.Allow(at(time.Hour)) {
			allowed++
		}
	}
	if allowed != 3 {
		t.Fatalf("after an hour idle, %d of 100 simultaneous requests were allowed, want 3 (the capacity)", allowed)
	}
}
