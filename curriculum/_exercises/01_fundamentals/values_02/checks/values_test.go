package values

import "testing"

func TestCanVote(t *testing.T) {
	tests := []struct {
		age        int
		registered bool
		want       bool
	}{
		{18, true, true},
		{40, true, true},
		{17, true, false},
		{18, false, false},
		{0, false, false},
	}
	for _, tt := range tests {
		if got := CanVote(tt.age, tt.registered); got != tt.want {
			t.Errorf("CanVote(%d, %v) = %v, want %v", tt.age, tt.registered, got, tt.want)
		}
	}
}

func TestDiffer(t *testing.T) {
	for _, tt := range []struct{ a, b, want bool }{
		{false, false, false},
		{true, false, true},
		{false, true, true},
		{true, true, false},
	} {
		if got := Differ(tt.a, tt.b); got != tt.want {
			t.Errorf("Differ(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestIsEven(t *testing.T) {
	for _, tt := range []struct {
		n    int
		want bool
	}{
		{0, true}, {2, true}, {7, false}, {-4, true}, {-3, false}, {1 << 40, true},
	} {
		if got := IsEven(tt.n); got != tt.want {
			t.Errorf("IsEven(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
