package countdown

import (
	"slices"
	"testing"
)

func TestCountdown(t *testing.T) {
	if got := slices.Collect(Countdown(3)); !slices.Equal(got, []int{3, 2, 1}) {
		t.Errorf("Countdown(3) yielded %v, want [3 2 1]", got)
	}
	if got := slices.Collect(Countdown(1)); !slices.Equal(got, []int{1}) {
		t.Errorf("Countdown(1) yielded %v, want [1]", got)
	}
	for _, n := range []int{0, -4} {
		if got := slices.Collect(Countdown(n)); len(got) != 0 {
			t.Errorf("Countdown(%d) yielded %v, want nothing", n, got)
		}
	}
}

func TestCountdownStopsOnBreak(t *testing.T) {
	var seen []int
	for i := range Countdown(10) {
		seen = append(seen, i)
		if i == 8 {
			break
		}
	}
	if !slices.Equal(seen, []int{10, 9, 8}) {
		t.Errorf("loop with break saw %v, want [10 9 8]", seen)
	}
}

func TestCountdownIsReusable(t *testing.T) {
	seq := Countdown(2)
	a, b := slices.Collect(seq), slices.Collect(seq)
	if !slices.Equal(a, b) || len(a) != 2 {
		t.Errorf("ranging twice gave %v and %v; an iter.Seq should start over each time", a, b)
	}
}
