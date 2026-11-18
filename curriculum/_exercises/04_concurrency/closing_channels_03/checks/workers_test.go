package workers

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestStopStopsEveryWorker(t *testing.T) {
	var calls atomic.Int64
	var seen [4]atomic.Bool
	stop := Start(4, func(id int) {
		seen[id].Store(true)
		time.Sleep(time.Millisecond)
		calls.Add(1)
	})
	time.Sleep(30 * time.Millisecond)

	done := make(chan struct{})
	go func() { stop(); close(done) }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("stop() did not return: not every worker was told to stop")
	}
	for id := range seen {
		if !seen[id].Load() {
			t.Errorf("worker %d never ran", id)
		}
	}
	after := calls.Load()
	time.Sleep(30 * time.Millisecond)
	if n := calls.Load(); n != after {
		t.Fatalf("work was called %d more times after stop returned", n-after)
	}
}
