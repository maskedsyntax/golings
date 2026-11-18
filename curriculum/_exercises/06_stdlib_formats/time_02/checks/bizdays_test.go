package bizdays

import (
	"testing"
	"time"
)

// March 2024: the 1st is a Friday.
func day(d int) time.Time { return time.Date(2024, 3, d, 9, 30, 0, 0, time.UTC) }

func TestAddBusinessDays(t *testing.T) {
	for _, tt := range []struct {
		from time.Time
		n    int
		want time.Time
	}{
		{day(1), 1, day(4)},   // Fri + 1 = Mon
		{day(6), 2, day(8)},   // Wed + 2 = Fri
		{day(7), 5, day(14)},  // Thu + 5 = next Thu
		{day(2), 1, day(4)},   // Sat + 1 = Mon
		{day(2), 0, day(2)},   // n == 0 leaves t alone
		{day(1), 10, day(15)}, // two full weeks
		{time.Date(2024, 2, 28, 9, 30, 0, 0, time.UTC), 2, day(1)}, // through Feb 29
	} {
		got := AddBusinessDays(tt.from, tt.n)
		if !got.Equal(tt.want) {
			t.Errorf("AddBusinessDays(%s, %d) = %s, want %s",
				tt.from.Format("Mon Jan 2"), tt.n, got.Format("Mon Jan 2 15:04"), tt.want.Format("Mon Jan 2 15:04"))
		}
	}
}
