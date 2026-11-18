package grade

import "testing"

func TestGrade(t *testing.T) {
	for _, tt := range []struct {
		score int
		want  string
	}{
		{100, "A"}, {90, "A"}, {89, "B"}, {80, "B"}, {79, "C"}, {70, "C"},
		{69, "D"}, {60, "D"}, {59, "F"}, {0, "F"}, {-1, "invalid"}, {101, "invalid"},
	} {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}
