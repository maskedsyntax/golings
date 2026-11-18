package try

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

func TestTrySend(t *testing.T) {
	ch := make(chan string, 1)
	var ok bool
	within(t, "TrySend", func() { ok = TrySend(ch, "hi") })
	if !ok || len(ch) != 1 {
		t.Fatalf("TrySend into an empty buffer = %v (len %d); want true", ok, len(ch))
	}
	within(t, "TrySend on a full channel", func() { ok = TrySend(ch, "again") })
	if ok {
		t.Fatal("TrySend on a full channel = true; want false")
	}
	within(t, "TrySend on an unbuffered channel", func() { ok = TrySend(make(chan string), "x") })
	if ok {
		t.Fatal("TrySend with no receiver = true; want false")
	}
}

func TestTryReceive(t *testing.T) {
	ch := make(chan string, 1)
	var got string
	var ok bool
	within(t, "TryReceive on an empty channel", func() { got, ok = TryReceive(ch) })
	if ok || got != "" {
		t.Fatalf("TryReceive(empty) = %q, %v; want \"\", false", got, ok)
	}
	ch <- "ready"
	within(t, "TryReceive", func() { got, ok = TryReceive(ch) })
	if !ok || got != "ready" {
		t.Fatalf("TryReceive = %q, %v; want \"ready\", true", got, ok)
	}
}
