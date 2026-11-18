package greet

import (
	"testing"
	"time"
)

func TestGreet(t *testing.T) {
	got := make(chan (<-chan string), 1)
	go func() { got <- Greet("Gopher") }()
	var ch <-chan string
	select {
	case ch = <-got:
	case <-time.After(2 * time.Second):
		t.Fatal("Greet did not return: it is stuck sending on a channel nobody can receive from yet")
	}
	select {
	case msg := <-ch:
		if msg != "Hello, Gopher!" {
			t.Fatalf("received %q, want %q", msg, "Hello, Gopher!")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("nothing was ever sent on the returned channel")
	}
}
