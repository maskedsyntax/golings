package relay

import (
	"runtime"
	"slices"
	"testing"
	"time"
)

func TestRelayRunsInOrder(t *testing.T) {
	var log []int
	steps := make([]func(), 8)
	for i := range steps {
		steps[i] = func() {
			if i == 0 {
				time.Sleep(20 * time.Millisecond) // a slow first runner
			}
			log = append(log, i)
		}
	}
	Relay(steps...)
	if want := []int{0, 1, 2, 3, 4, 5, 6, 7}; !slices.Equal(log, want) {
		t.Fatalf("steps ran as %v, want %v", log, want)
	}
}

func TestRelayUsesGoroutines(t *testing.T) {
	main := goid()
	same := false
	Relay(func() { same = goid() == main }, func() {})
	if same {
		t.Fatal("a step ran in Relay's own goroutine; run each step in a new goroutine")
	}
}

func TestRelayEmpty(t *testing.T) {
	done := make(chan struct{})
	go func() { Relay(); close(done) }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Relay() with no steps blocked")
	}
}

// goid returns an identifier for the current goroutine (test helper only).
func goid() string {
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]
	// "goroutine 18 [running]:..."
	for i := len("goroutine "); i < len(buf); i++ {
		if buf[i] == ' ' {
			return string(buf[len("goroutine "):i])
		}
	}
	return string(buf)
}
