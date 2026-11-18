package timeout

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

func TestFastFetch(t *testing.T) {
	var got string
	var err error
	within(t, "WithTimeout", func() { got, err = WithTimeout(func() string { return "data" }, time.Second) })
	if got != "data" || err != nil {
		t.Fatalf("WithTimeout(fast) = %q, %v; want \"data\", nil", got, err)
	}
}

func TestSlowFetchTimesOut(t *testing.T) {
	release := make(chan struct{})
	defer close(release)
	var got string
	var err error
	start := time.Now()
	within(t, "WithTimeout(slow fetch)", func() {
		got, err = WithTimeout(func() string { <-release; return "late" }, 50*time.Millisecond)
	})
	if !errors.Is(err, ErrTimeout) || got != "" {
		t.Fatalf("WithTimeout(slow) = %q, %v; want \"\", ErrTimeout", got, err)
	}
	if d := time.Since(start); d > time.Second {
		t.Fatalf("timing out took %v; it should take about 50ms", d)
	}
}

func TestNoLeak(t *testing.T) {
	before := runtime.NumGoroutine()
	release := make(chan struct{})
	within(t, "WithTimeout(slow fetch)", func() {
		WithTimeout(func() string { <-release; return "late" }, 10*time.Millisecond)
	})
	close(release) // fetch finishes now; its goroutine must be able to exit
	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > before {
		if time.Now().After(deadline) {
			t.Fatal("a goroutine leaked: after timing out, nobody receives fetch's result, so its send blocks forever")
		}
		time.Sleep(10 * time.Millisecond)
	}
}
