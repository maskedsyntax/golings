package layouts

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	tm := time.Date(2024, 3, 5, 18, 0, 0, 0, time.UTC)
	if got, want := FormatDate(tm), "2024-03-05"; got != want {
		t.Errorf("FormatDate = %q, want %q", got, want)
	}
	tm = time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)
	if got, want := FormatDate(tm), "1999-12-31"; got != want {
		t.Errorf("FormatDate = %q, want %q", got, want)
	}
}

func TestParseDate(t *testing.T) {
	got, err := ParseDate("2024-03-05")
	if err != nil {
		t.Fatalf("ParseDate returned an error: %v", err)
	}
	if want := time.Date(2024, 3, 5, 0, 0, 0, 0, time.UTC); !got.Equal(want) {
		t.Errorf("ParseDate = %s, want %s", got, want)
	}
	for _, bad := range []string{"2024-13-01", "05/03/2024", "2024-03-05 10:00"} {
		if _, err := ParseDate(bad); err == nil {
			t.Errorf("ParseDate(%q) returned no error", bad)
		}
	}
}

func TestClock(t *testing.T) {
	for _, tt := range []struct {
		h, m int
		want string
	}{
		{15, 4, "3:04 PM"},
		{9, 5, "9:05 AM"},
		{0, 0, "12:00 AM"},
		{12, 30, "12:30 PM"},
	} {
		tm := time.Date(2024, 3, 5, tt.h, tt.m, 0, 0, time.UTC)
		if got := Clock(tm); got != tt.want {
			t.Errorf("Clock(%02d:%02d) = %q, want %q", tt.h, tt.m, got, tt.want)
		}
	}
}
