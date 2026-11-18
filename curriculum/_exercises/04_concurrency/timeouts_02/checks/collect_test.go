package collect

import (
	"testing"
	"time"
)

// within fails the test if f does not return within 2 seconds.
func within(t *testing.T, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2s", what)
	}
}

func ticker(stop <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- i:
				time.Sleep(5 * time.Millisecond)
			case <-stop:
				return
			}
		}
	}()
	return ch
}

func TestCollectForStopsDespiteSteadyStream(t *testing.T) {
	stop := make(chan struct{})
	defer close(stop)
	var got []int
	start := time.Now()
	within(t, "CollectFor (the timer restarts with every value)", func() { got = CollectFor(ticker(stop), 100*time.Millisecond) })
	if d := time.Since(start); d > 700*time.Millisecond {
		t.Fatalf("CollectFor(100ms) took %v", d)
	}
	if len(got) < 2 {
		t.Fatalf("collected %d values in 100ms; expected several", len(got))
	}
}

func TestCollectForQuietChannel(t *testing.T) {
	var got []int
	within(t, "CollectFor", func() { got = CollectFor(make(chan int), 20*time.Millisecond) })
	if len(got) != 0 {
		t.Fatalf("CollectFor on a silent channel = %v", got)
	}
}
