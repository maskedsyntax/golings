package collect

import (
	"slices"
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

func TestCollect(t *testing.T) {
	data := make(chan int)
	quit := make(chan struct{})
	go func() {
		for i := 1; i <= 4; i++ {
			data <- i // unbuffered: each send completes only once Collect has it
		}
		close(quit)
	}()
	var got []int
	within(t, "Collect", func() { got = Collect(data, quit) })
	if want := []int{1, 2, 3, 4}; !slices.Equal(got, want) {
		t.Fatalf("Collect = %v, want %v", got, want)
	}
}

func TestCollectQuitImmediately(t *testing.T) {
	quit := make(chan struct{})
	close(quit)
	var got []int
	within(t, "Collect", func() { got = Collect(make(chan int), quit) })
	if len(got) != 0 {
		t.Fatalf("Collect = %v, want nothing", got)
	}
}
