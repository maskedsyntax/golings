package ttl

import (
	"testing"
	"time"
)

func TestTTL(t *testing.T) {
	for _, tt := range []struct {
		exp  int64
		now  time.Time
		want time.Duration
	}{
		{1090, time.Unix(1000, 0), 90 * time.Second},
		{1010, time.Unix(1000, 500_000_000), 9500 * time.Millisecond},
		{1000, time.Unix(1000, 0), 0},
		{900, time.Unix(1000, 0), 0},
		{1_700_000_000 + 3600, time.Unix(1_700_000_000, 0), time.Hour},
	} {
		if got := TTL(tt.exp, tt.now); got != tt.want {
			t.Errorf("TTL(%d, %d) = %v, want %v", tt.exp, tt.now.Unix(), got, tt.want)
		}
	}
}
