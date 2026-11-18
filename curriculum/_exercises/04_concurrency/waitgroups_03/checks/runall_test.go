package runall

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

var (
	errA = errors.New("a failed")
	errC = errors.New("c failed")
)

func TestRunAllIsConcurrent(t *testing.T) {
	// Every task waits until all three have started.
	var started atomic.Int32
	all := make(chan struct{})
	deadline := make(chan struct{})
	time.AfterFunc(time.Second, func() { close(deadline) })
	task := func(err error) func() error {
		return func() error {
			if started.Add(1) == 3 {
				close(all)
			}
			select {
			case <-all:
				return err
			case <-deadline:
				return errors.New("timed out waiting for the other tasks: they are not running concurrently")
			}
		}
	}
	err := RunAll(task(errA), task(nil), task(errC))
	if err == nil {
		t.Fatal("RunAll = nil, want the joined errors")
	}
	if want := "a failed\nc failed"; err.Error() != want {
		t.Fatalf("RunAll error = %q, want %q", err.Error(), want)
	}
	if !errors.Is(err, errA) || !errors.Is(err, errC) {
		t.Fatal("the joined error should match both errA and errC with errors.Is")
	}
}

func TestRunAllSuccess(t *testing.T) {
	ok := func() error { return nil }
	if err := RunAll(ok, ok); err != nil {
		t.Fatalf("RunAll = %v, want nil", err)
	}
	if err := RunAll(); err != nil {
		t.Fatalf("RunAll() = %v, want nil", err)
	}
}
