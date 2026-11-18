package schedule

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestScheduleRuns(t *testing.T) {
	fired := make(chan struct{})
	Schedule(20*time.Millisecond, func() { close(fired) })
	select {
	case <-fired:
	case <-time.After(2 * time.Second):
		t.Fatal("f never ran")
	}
}

func TestCancelBeforeFiring(t *testing.T) {
	var ran atomic.Bool
	cancel := Schedule(50*time.Millisecond, func() { ran.Store(true) })
	if !cancel() {
		t.Fatal("cancel() before the timer fired = false, want true")
	}
	time.Sleep(120 * time.Millisecond)
	if ran.Load() {
		t.Fatal("f ran even though it was cancelled")
	}
}

func TestCancelAfterFiring(t *testing.T) {
	fired := make(chan struct{})
	cancel := Schedule(5*time.Millisecond, func() { close(fired) })
	select {
	case <-fired:
	case <-time.After(2 * time.Second):
		t.Fatal("f never ran")
	}
	if cancel() {
		t.Fatal("cancel() after f ran = true, want false")
	}
}
