package weekend

import "testing"

func TestIsWeekend(t *testing.T) {
	for _, tt := range []struct {
		day  string
		want bool
	}{
		{"Saturday", true}, {"Sunday", true}, {"Monday", false}, {"Friday", false},
		{"saturday", false}, {"", false},
	} {
		if got := IsWeekend(tt.day); got != tt.want {
			t.Errorf("IsWeekend(%q) = %v, want %v", tt.day, got, tt.want)
		}
	}
}
