package store

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

func TestDelete(t *testing.T) {
	s := NewStore()
	s.Set("a", 1)
	s.Set("b", 2)
	within(t, "Delete", func() { s.Delete("a") })
	var a, b int
	within(t, "Get after Delete", func() { a, b = s.Get("a"), s.Get("b") })
	if a != 0 || b != 2 {
		t.Fatalf("after Delete(\"a\"): a = %d, b = %d; want 0 and 2", a, b)
	}
}
