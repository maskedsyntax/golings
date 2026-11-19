package wait

import (
	"context"
	"errors"
	"testing"
	"time"
)

// within fails the test if f doesn't return in time, instead of hanging.
func within(t *testing.T, what string, d time.Duration, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	select {
	case <-done:
	case <-time.After(d):
		t.Fatalf("%s did not finish within %v", what, d)
	}
}

func TestWaitFull(t *testing.T) {
	start := time.Now()
	if err := Wait(context.Background(), 30*time.Millisecond); err != nil {
		t.Fatalf("Wait returned %v, want nil", err)
	}
	if d := time.Since(start); d < 30*time.Millisecond {
		t.Fatalf("Wait returned after %v; it should wait the full 30ms", d)
	}
}

func TestWaitCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(20*time.Millisecond, cancel)
	within(t, "Wait after cancel", time.Second, func() {
		if err := Wait(ctx, 10*time.Second); !errors.Is(err, context.Canceled) {
			t.Errorf("Wait returned %v, want context.Canceled", err)
		}
	})
}

func TestWaitDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	within(t, "Wait past the deadline", time.Second, func() {
		if err := Wait(ctx, 10*time.Second); !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("Wait returned %v, want context.DeadlineExceeded", err)
		}
	})
}
