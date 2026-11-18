package pipeline

import (
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	done := make(chan int, 1)
	go func() { done <- Count(Filter(Generate(1, 2, 3, 4, 5, 6, 8), isEven)) }()
	select {
	case got := <-done:
		if got != 4 {
			t.Fatalf("Count = %d, want 4", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("the pipeline never finished: some stage doesn't close its output")
	}
}
