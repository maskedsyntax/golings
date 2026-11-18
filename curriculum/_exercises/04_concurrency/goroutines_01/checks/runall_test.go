package runall

import (
	"sync/atomic"
	"testing"
	"time"
)

// barrierTasks returns n tasks that each wait until all n have started.
// Run one at a time, the first task waits forever (until the deadline).
func barrierTasks(n int, deadline <-chan struct{}) ([]func(), *atomic.Int32, *atomic.Bool) {
	var arrived atomic.Int32
	var timedOut atomic.Bool
	all := make(chan struct{})
	tasks := make([]func(), n)
	for i := range tasks {
		tasks[i] = func() {
			if arrived.Add(1) == int32(n) {
				close(all)
			}
			select {
			case <-all:
			case <-deadline:
				timedOut.Store(true)
			}
		}
	}
	return tasks, &arrived, &timedOut
}

func TestRunAllRunsTasksConcurrently(t *testing.T) {
	deadline := make(chan struct{})
	time.AfterFunc(time.Second, func() { close(deadline) })
	tasks, arrived, timedOut := barrierTasks(5, deadline)
	RunAll(tasks...)
	if arrived.Load() != 5 {
		t.Fatalf("%d of 5 tasks ran", arrived.Load())
	}
	if timedOut.Load() {
		t.Fatal("the tasks did not run at the same time: the first one waited for the others, which hadn't started. Start each task with go.")
	}
}

func TestRunAllWaitsForTasks(t *testing.T) {
	var done atomic.Int32
	slow := func() { time.Sleep(20 * time.Millisecond); done.Add(1) }
	RunAll(slow, slow, slow)
	if done.Load() != 3 {
		t.Fatalf("RunAll returned before its tasks finished (%d of 3 done)", done.Load())
	}
}

func TestRunAllNoTasks(t *testing.T) {
	RunAll() // must not block
}
