package limiter

import (
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

func TestBurstThenLimited(t *testing.T) {
	l := NewLimiter(3, 60*time.Millisecond)
	defer l.Stop()
	start := time.Now()
	within(t, "the first 3 Waits", func() { l.Wait(); l.Wait(); l.Wait() })
	if d := time.Since(start); d > 30*time.Millisecond {
		t.Fatalf("the first burst of 3 took %v; the limiter should start full", d)
	}
	start = time.Now()
	within(t, "the 4th Wait", l.Wait)
	if d := time.Since(start); d < 30*time.Millisecond {
		t.Fatalf("the 4th Wait returned after %v; it should wait for a refill (about 60ms)", d)
	}
}

func TestIdleDoesNotOverfill(t *testing.T) {
	l := NewLimiter(2, 20*time.Millisecond)
	defer l.Stop()
	within(t, "draining the initial burst", func() { l.Wait(); l.Wait() })
	time.Sleep(200 * time.Millisecond) // about 10 refills, but only 2 fit
	within(t, "2 Waits after idling", func() { l.Wait(); l.Wait() })
	start := time.Now()
	within(t, "the 3rd Wait after idling", l.Wait)
	if d := time.Since(start); d < 8*time.Millisecond {
		t.Fatalf("a 3rd Wait after idling returned in %v; the limiter must not hold more than burst tokens", d)
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
