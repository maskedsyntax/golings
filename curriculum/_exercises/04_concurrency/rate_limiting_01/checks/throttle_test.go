package throttle

import (
	"slices"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	var got []string
	var times []time.Time
	start := time.Now()
	Throttle([]string{"a", "b", "c", "d", "e"}, 25*time.Millisecond, func(r string) {
		got = append(got, r)
		times = append(times, time.Now())
	})
	if !slices.Equal(got, []string{"a", "b", "c", "d", "e"}) {
		t.Fatalf("handled %q, want a..e in order", got)
	}
	if d := time.Since(start); d < 110*time.Millisecond {
		t.Fatalf("5 requests at one per 25ms took only %v", d)
	}
	for i := 1; i < len(times); i++ {
		if gap := times[i].Sub(times[i-1]); gap < 10*time.Millisecond {
			t.Fatalf("requests %d and %d were only %v apart", i-1, i, gap)
		}
	}
}
