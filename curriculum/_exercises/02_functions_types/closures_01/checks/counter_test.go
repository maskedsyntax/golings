package counter

import "testing"

func TestCounterCounts(t *testing.T) {
	next := Counter()
	for want := 1; want <= 5; want++ {
		if got := next(); got != want {
			t.Fatalf("call %d returned %d, want %d", want, got, want)
		}
	}
}

func TestCountersAreIndependent(t *testing.T) {
	a := Counter()
	a()
	a()
	b := Counter()
	if got := b(); got != 1 {
		t.Errorf("a new counter's first call returned %d, want 1", got)
	}
	if got := a(); got != 3 {
		t.Errorf("the first counter's third call returned %d, want 3", got)
	}
}
