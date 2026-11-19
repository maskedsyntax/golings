package first

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

func answer(s string, after time.Duration) func(context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		select {
		case <-time.After(after):
			return s, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func failure(err error) func(context.Context) (string, error) {
	return func(context.Context) (string, error) { return "", err }
}

func TestFirstSuccessWins(t *testing.T) {
	cancelled := make(chan struct{})
	slow := func(ctx context.Context) (string, error) {
		select {
		case <-ctx.Done():
			close(cancelled)
			return "", ctx.Err()
		case <-time.After(5 * time.Second):
			return "slow", nil
		}
	}
	within(t, "FirstResult", 2*time.Second, func() {
		got, err := FirstResult(context.Background(), slow, failure(errors.New("down")), answer("fast", 20*time.Millisecond))
		if err != nil || got != "fast" {
			t.Errorf("FirstResult = %q, %v; want \"fast\", nil", got, err)
		}
	})
	select {
	case <-cancelled:
	case <-time.After(time.Second):
		t.Fatal("the slow function's context was never cancelled after the winner returned")
	}
}

func TestAllFail(t *testing.T) {
	e1, e2 := errors.New("one"), errors.New("two")
	within(t, "FirstResult", 2*time.Second, func() {
		_, err := FirstResult(context.Background(), failure(e1), failure(e2))
		if !errors.Is(err, e1) || !errors.Is(err, e2) {
			t.Errorf("FirstResult error = %v; want one matching both failures", err)
		}
	})
}

func TestParentCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(20*time.Millisecond, cancel)
	within(t, "FirstResult", 2*time.Second, func() {
		_, err := FirstResult(ctx, answer("never", 10*time.Second))
		if !errors.Is(err, context.Canceled) {
			t.Errorf("FirstResult error = %v, want context.Canceled", err)
		}
	})
}

func TestNoFunctions(t *testing.T) {
	within(t, "FirstResult", time.Second, func() {
		if _, err := FirstResult(context.Background()); err == nil {
			t.Error("FirstResult() with no functions returned no error")
		}
	})
}
