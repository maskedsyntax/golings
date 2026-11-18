package logtime

import (
	"testing"
	"time"
)

func TestParseLogTime(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want time.Time
	}{
		{"2024-03-05T10:00:00+02:00", time.Date(2024, 3, 5, 8, 0, 0, 0, time.UTC)},
		{"2024-03-05T08:00:00Z", time.Date(2024, 3, 5, 8, 0, 0, 0, time.UTC)},
		{"2024-03-05T08:00:00.25Z", time.Date(2024, 3, 5, 8, 0, 0, 250_000_000, time.UTC)},
		{"05/Mar/2024:10:00:00 +0200", time.Date(2024, 3, 5, 8, 0, 0, 0, time.UTC)},
		{"31/Dec/2023:23:30:00 -0100", time.Date(2024, 1, 1, 0, 30, 0, 0, time.UTC)},
	} {
		got, err := ParseLogTime(tt.in)
		if err != nil {
			t.Errorf("ParseLogTime(%q) returned an error: %v", tt.in, err)
			continue
		}
		if !got.Equal(tt.want) {
			t.Errorf("ParseLogTime(%q) = %s, want %s", tt.in, got, tt.want)
		}
		if got.Location() != time.UTC {
			t.Errorf("ParseLogTime(%q) is in %s, want UTC", tt.in, got.Location())
		}
	}
}

func TestParseLogTimeErrors(t *testing.T) {
	for _, in := range []string{"2024-03-05", "05/Mar/2024 10:00:00", "yesterday", ""} {
		if _, err := ParseLogTime(in); err == nil {
			t.Errorf("ParseLogTime(%q) returned no error", in)
		}
	}
}
