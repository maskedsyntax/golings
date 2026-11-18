package epoch

import (
	"testing"
	"time"
)

func TestToUnixMillis(t *testing.T) {
	tm := time.Date(2024, 3, 15, 10, 30, 0, 123_000_000, time.UTC)
	if got, want := ToUnixMillis(tm), int64(1710498600123); got != want {
		t.Errorf("ToUnixMillis(%s) = %d, want %d", tm, got, want)
	}
	before := time.Date(1969, 12, 31, 23, 59, 58, 500_000_000, time.UTC)
	if got, want := ToUnixMillis(before), int64(-1500); got != want {
		t.Errorf("ToUnixMillis(%s) = %d, want %d", before, got, want)
	}
}

func TestFromUnixMillis(t *testing.T) {
	for _, tt := range []struct {
		ms   int64
		want time.Time
	}{
		{1710498600123, time.Date(2024, 3, 15, 10, 30, 0, 123_000_000, time.UTC)},
		{0, time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
		{-1500, time.Date(1969, 12, 31, 23, 59, 58, 500_000_000, time.UTC)},
	} {
		got := FromUnixMillis(tt.ms)
		if !got.Equal(tt.want) {
			t.Errorf("FromUnixMillis(%d) = %s, want %s", tt.ms, got, tt.want)
		}
		if got.Location() != time.UTC {
			t.Errorf("FromUnixMillis(%d) is in %s, want UTC", tt.ms, got.Location())
		}
	}
}
