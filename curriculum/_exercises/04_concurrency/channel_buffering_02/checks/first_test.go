package first

import (
	"runtime"
	"testing"
	"time"
)

func TestFirstReturnsFastest(t *testing.T) {
	release := make(chan struct{})
	slow := func() string { <-release; return "slow" }
	fast := func() string { return "fast" }
	got := First(slow, fast, slow)
	close(release)
	if got != "fast" {
		t.Fatalf("First = %q, want %q", got, "fast")
	}
}

func TestFirstDoesNotLeak(t *testing.T) {
	before := runtime.NumGoroutine()
	release := make(chan struct{})
	slow := func() string { <-release; return "slow" }
	fast := func() string { return "fast" }
	First(fast, slow, slow, slow, slow)
	close(release) // the slow sources finish now

	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > before {
		if time.Now().After(deadline) {
			t.Fatalf("%d goroutines leaked: the losing sources are stuck sending their answers", runtime.NumGoroutine()-before)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
