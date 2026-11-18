package limiter

import (
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

func TestBurstThenLimited(t *testing.T) {
	l := NewLimiter(3, 60*time.Millisecond)
	defer l.Stop()
	start := time.Now()
	within(t, "the first 3 Waits", func() { l.Wait(); l.Wait(); l.Wait() })
	if d := time.Since(start); d > 40*time.Millisecond {
		t.Fatalf("the first burst of 3 took %v; the limiter should start full", d)
	}
	start = time.Now()
	within(t, "the 4th Wait", l.Wait)
	if d := time.Since(start); d < 15*time.Millisecond {
		t.Fatalf("the 4th Wait returned after %v; it should wait for a refill (about 60ms)", d)
	}
}

func TestIdleDoesNotOverfill(t *testing.T) {
	l := NewLimiter(2, 100*time.Millisecond)
	defer l.Stop()
	within(t, "draining the initial burst", func() { l.Wait(); l.Wait() })
	time.Sleep(650 * time.Millisecond) // about 6 refills, but only 2 fit

	// Count how many Waits get through in the next 40ms. A full bucket
	// allows 2, plus at most 1 more if a refill happens to land meanwhile.
	var passed atomic.Int32
	go func() {
		for {
			l.Wait()
			passed.Add(1)
		}
	}()
	time.Sleep(40 * time.Millisecond)
	if n := passed.Load(); n > 3 {
		t.Fatalf("%d Waits returned at once after idling; the limiter must not hold more than burst (2) tokens", n)
	} else if n < 2 {
		t.Fatalf("only %d Wait(s) returned after idling; the bucket should have refilled to 2", n)
	}
}

func TestStopEndsRefill(t *testing.T) {
	before := runtime.NumGoroutine()
	l := NewLimiter(1, 5*time.Millisecond)
	l.Stop()
	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > before {
		if time.Now().After(deadline) {
			t.Fatal("the refill goroutine is still running after Stop")
		}
		time.Sleep(5 * time.Millisecond)
	}
}
