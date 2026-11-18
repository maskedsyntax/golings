package poll

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

func TestPollWaitsBetweenChecks(t *testing.T) {
	calls := 0
	check := func() bool { calls++; return calls == 4 }
	var n int
	start := time.Now()
	within(t, "Poll", func() { n = Poll(20*time.Millisecond, check) })
	elapsed := time.Since(start)
	if n != 4 || calls != 4 {
		t.Fatalf("Poll = %d after %d calls, want 4", n, calls)
	}
	if elapsed < 70*time.Millisecond {
		t.Fatalf("4 checks at a 20ms interval took only %v; wait for a tick before each check", elapsed)
	}
}

func TestPollImmediateSuccess(t *testing.T) {
	var n int
	within(t, "Poll", func() { n = Poll(10*time.Millisecond, func() bool { return true }) })
	if n != 1 {
		t.Fatalf("Poll = %d, want 1", n)
	}
}
