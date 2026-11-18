package queue

import (
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	msgs := []string{"buffered", "channel", "values"}
	done := make(chan chan string, 1)
	go func() { done <- Queue(msgs...) }()
	var ch chan string
	select {
	case ch = <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Queue blocked: a send had no room in the buffer and no receiver")
	}
	if len(ch) != len(msgs) {
		t.Fatalf("the channel holds %d values, want %d", len(ch), len(msgs))
	}
	for _, want := range msgs {
		if got := <-ch; got != want {
			t.Fatalf("received %q, want %q", got, want)
		}
	}
}

func TestQueueEmpty(t *testing.T) {
	if ch := Queue(); ch == nil || len(ch) != 0 {
		t.Fatalf("Queue() = %v, want an empty channel", ch)
	}
}
