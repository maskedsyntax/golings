package process

import (
	"errors"
	"fmt"
	"runtime"
	"sync/atomic"
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

func items(n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = i
	}
	return out
}

func TestAllSucceed(t *testing.T) {
	var calls atomic.Int32
	var err error
	within(t, "ProcessAll", func() {
		err = ProcessAll(items(50), 4, func(int) error { calls.Add(1); return nil })
	})
	if err != nil || calls.Load() != 50 {
		t.Fatalf("ProcessAll = %v after %d calls; want nil after 50", err, calls.Load())
	}
}

var errBad = errors.New("bad item")

func TestStopsOnFirstError(t *testing.T) {
	var calls atomic.Int32
	var err error
	within(t, "ProcessAll", func() {
		err = ProcessAll(items(200), 4, func(i int) error {
			calls.Add(1)
			time.Sleep(time.Millisecond)
			if i == 10 {
				return fmt.Errorf("item %d: %w", i, errBad)
			}
			return nil
		})
	})
	if !errors.Is(err, errBad) {
		t.Fatalf("ProcessAll = %v, want the error from item 10", err)
	}
	if n := calls.Load(); n > 60 {
		t.Fatalf("f was called %d times; dispatching should stop soon after the error at item 10", n)
	}
}

func TestNoLeakedGoroutines(t *testing.T) {
	before := runtime.NumGoroutine()
	within(t, "ProcessAll", func() {
		ProcessAll(items(100), 8, func(i int) error {
			if i%7 == 3 {
				return errBad
			}
			return nil
		})
	})
	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > before {
		if time.Now().After(deadline) {
			t.Fatalf("%d goroutines still running after ProcessAll returned", runtime.NumGoroutine()-before)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
