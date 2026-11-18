package collatz

import "testing"

func TestSteps(t *testing.T) {
	for _, tt := range []struct{ n, want int }{
		{1, 0}, {2, 1}, {6, 8}, {7, 16}, {27, 111}, {0, -1}, {-8, -1},
	} {
		if got := Steps(tt.n); got != tt.want {
			t.Errorf("Steps(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
