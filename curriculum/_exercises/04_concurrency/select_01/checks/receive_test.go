package receive

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

func TestReceiveFromWhicheverIsReady(t *testing.T) {
	a, b := make(chan string), make(chan string)
	go func() {
		b <- "b1"
		b <- "b2"
		b <- "b3"
		a <- "a1" // a only sends once b is done
	}()
	var got []string
	within(t, "Receive", func() { got = Receive(a, b, 4) })
	if want := []string{"b1", "b2", "b3", "a1"}; !slices.Equal(got, want) {
		t.Fatalf("Receive = %q, want %q", got, want)
	}
}

func TestReceiveOnlyA(t *testing.T) {
	a, b := make(chan string), make(chan string)
	go func() { a <- "x"; a <- "y" }()
	var got []string
	within(t, "Receive", func() { got = Receive(a, b, 2) })
	if want := []string{"x", "y"}; !slices.Equal(got, want) {
		t.Fatalf("Receive = %q, want %q", got, want)
	}
}
