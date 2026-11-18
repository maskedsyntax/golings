package start

import (
	"testing"
	"time"
)

func TestDoneWaitsForTask(t *testing.T) {
	release := make(chan struct{})
	finished := false
	done := Start(func() {
		<-release
		finished = true
	})
	select {
	case <-done:
		t.Fatal("done was closed while the task was still running")
	case <-time.After(50 * time.Millisecond):
	}
	close(release)
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("done was never closed after the task finished")
	}
	if !finished {
		t.Fatal("done was closed before the task finished")
	}
}
