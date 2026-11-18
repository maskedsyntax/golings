package sum

import (
	"testing"
	"time"
)

// within fails the test if f does not return within 2 seconds.
func within(t *testing.T, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2s", what)
	}
}

func TestSum(t *testing.T) {
	data := make(chan int)
	done := make(chan struct{})
	go func() {
		for _, v := range []int{5, 10, 20} {
			data <- v
		}
		close(done)
	}()
	var got int
	within(t, "Sum (break only leaves the select, not the loop)", func() { got = Sum(data, done) })
	if got != 35 {
		t.Fatalf("Sum = %d, want 35", got)
	}
}
