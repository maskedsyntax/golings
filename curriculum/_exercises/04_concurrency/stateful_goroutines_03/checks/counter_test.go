package counter

import (
	"errors"
	"runtime"
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

func TestIncThenClose(t *testing.T) {
	c := NewCounter()
	for want := 1; want <= 3; want++ {
		var got int
		var err error
		within(t, "Inc", func() { got, err = c.Inc() })
		if got != want || err != nil {
			t.Fatalf("Inc = %d, %v; want %d, nil", got, err, want)
		}
	}
	within(t, "Close", c.Close)
	var err error
	within(t, "Inc after Close", func() { _, err = c.Inc() })
	if !errors.Is(err, ErrClosed) {
		t.Fatalf("Inc after Close returned %v, want ErrClosed", err)
	}
}

func TestCloseStopsOwner(t *testing.T) {
	before := runtime.NumGoroutine()
	c := NewCounter()
	within(t, "Close", c.Close)
	deadline := time.Now().Add(500 * time.Millisecond)
	for runtime.NumGoroutine() > before {
		if time.Now().After(deadline) {
			t.Fatalf("%d goroutine(s) still running after Close; the owner should exit", runtime.NumGoroutine()-before)
		}
		time.Sleep(5 * time.Millisecond)
	}
}
