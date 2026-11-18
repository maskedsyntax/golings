package variadic

import "testing"

func TestMax(t *testing.T) {
	if got := Max(3, 9, -2); got != 9 {
		t.Errorf("Max(3, 9, -2) = %d, want 9", got)
	}
	if got := Max(-7); got != -7 {
		t.Errorf("Max(-7) = %d, want -7", got)
	}
	if got := Max(-5, -3, -8); got != -3 {
		t.Errorf("Max(-5, -3, -8) = %d, want -3", got)
	}
	nums := []int{4, 12, 1}
	if got := Max(0, nums...); got != 12 {
		t.Errorf("Max(0, %v...) = %d, want 12", nums, got)
	}
}

func TestJoinNonEmpty(t *testing.T) {
	for _, tt := range []struct {
		sep   string
		parts []string
		want  string
	}{
		{", ", []string{"a", "", "b"}, "a, b"},
		{"-", []string{"x"}, "x"},
		{"/", []string{"", "usr", "", "bin", ""}, "usr/bin"},
		{"+", nil, ""},
		{"+", []string{"", ""}, ""},
	} {
		if got := JoinNonEmpty(tt.sep, tt.parts...); got != tt.want {
			t.Errorf("JoinNonEmpty(%q, %q...) = %q, want %q", tt.sep, tt.parts, got, tt.want)
		}
	}
}
