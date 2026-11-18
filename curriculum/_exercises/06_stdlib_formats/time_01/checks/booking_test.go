package booking

import (
	"testing"
	"time"
)

var plus2 = time.FixedZone("UTC+2", 2*60*60)

func at(h, m int) time.Time { return time.Date(2024, 3, 5, h, m, 0, 0, time.UTC) }

func TestSameMoment(t *testing.T) {
	if !SameMoment(at(12, 0), time.Date(2024, 3, 5, 14, 0, 0, 0, plus2)) {
		t.Error("12:00 UTC and 14:00 UTC+2 are the same moment")
	}
	if SameMoment(at(12, 0), time.Date(2024, 3, 5, 12, 0, 0, 0, plus2)) {
		t.Error("12:00 UTC and 12:00 UTC+2 are different moments")
	}
	if !SameMoment(at(9, 30), at(9, 30)) {
		t.Error("a time should be the same moment as itself")
	}
}

func TestOverlaps(t *testing.T) {
	for _, tt := range []struct {
		name string
		a, b Booking
		want bool
	}{
		{"touching", Booking{at(9, 0), at(10, 0)}, Booking{at(10, 0), at(11, 0)}, false},
		{"apart", Booking{at(9, 0), at(10, 0)}, Booking{at(11, 0), at(12, 0)}, false},
		{"partial", Booking{at(9, 0), at(10, 0)}, Booking{at(9, 30), at(11, 0)}, true},
		{"inside", Booking{at(9, 0), at(12, 0)}, Booking{at(10, 0), at(11, 0)}, true},
		{"other zone", Booking{at(9, 0), at(10, 0)},
			Booking{time.Date(2024, 3, 5, 11, 30, 0, 0, plus2), time.Date(2024, 3, 5, 13, 0, 0, 0, plus2)}, true},
		{"other zone touching", Booking{at(9, 0), at(10, 0)},
			Booking{time.Date(2024, 3, 5, 12, 0, 0, 0, plus2), time.Date(2024, 3, 5, 13, 0, 0, 0, plus2)}, false},
	} {
		if got := Overlaps(tt.a, tt.b); got != tt.want {
			t.Errorf("%s: Overlaps = %v, want %v", tt.name, got, tt.want)
		}
		if got := Overlaps(tt.b, tt.a); got != tt.want {
			t.Errorf("%s (swapped): Overlaps = %v, want %v", tt.name, got, tt.want)
		}
	}
}
