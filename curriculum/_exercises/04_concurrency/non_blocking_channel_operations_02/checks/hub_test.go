package hub

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

func TestPublishSkipsFullSubscribers(t *testing.T) {
	var h Hub
	stuck := h.Subscribe(1) // never read
	live := h.Subscribe(10)
	within(t, "Publish (a full subscriber blocked it)", func() {
		for v := range 3 {
			h.Publish(v)
		}
	})
	if len(live) != 3 {
		t.Fatalf("the live subscriber got %d values, want 3", len(live))
	}
	if len(stuck) != 1 || <-stuck != 0 {
		t.Fatal("the stuck subscriber should hold only the first value")
	}
	if h.Dropped != 2 {
		t.Fatalf("Dropped = %d, want 2", h.Dropped)
	}
}

func TestPublishDeliversWhenThereIsRoom(t *testing.T) {
	var h Hub
	a, b := h.Subscribe(2), h.Subscribe(2)
	within(t, "Publish", func() { h.Publish(7); h.Publish(8) })
	for _, s := range []<-chan int{a, b} {
		if x, y := <-s, <-s; x != 7 || y != 8 {
			t.Fatalf("subscriber got %d, %d; want 7, 8", x, y)
		}
	}
	if h.Dropped != 0 {
		t.Fatalf("Dropped = %d, want 0", h.Dropped)
	}
}
