package heartbeat

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestHeartbeatStops(t *testing.T) {
	var beats atomic.Int32
	stop := make(chan struct{})
	done := make(chan struct{})
	go func() {
		Heartbeat(10*time.Millisecond, stop, func() { beats.Add(1) })
		close(done)
	}()
	time.Sleep(75 * time.Millisecond)
	close(stop)
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Heartbeat did not return after stop was closed")
	}
	n := beats.Load()
	if n < 3 {
		t.Fatalf("only %d beats in 75ms at a 10ms interval", n)
	}
	time.Sleep(50 * time.Millisecond)
	if beats.Load() != n {
		t.Fatal("beat was called after Heartbeat returned")
	}
}
