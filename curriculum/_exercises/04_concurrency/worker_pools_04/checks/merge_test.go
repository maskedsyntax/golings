package merge

import (
	"slices"
	"testing"
	"time"
)

func source(from, to int, delay time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := from; i < to; i++ {
			time.Sleep(delay)
			ch <- i
		}
	}()
	return ch
}

func drain(t *testing.T, ch <-chan int) []int {
	t.Helper()
	if ch == nil {
		t.Fatal("Merge returned a nil channel")
	}
	var out []int
	timeout := time.After(3 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				slices.Sort(out)
				return out
			}
			out = append(out, v)
		case <-timeout:
			t.Fatalf("received %d values and then the output was never closed", len(out))
		}
	}
}

func TestMerge(t *testing.T) {
	got := drain(t, Merge(source(0, 10, time.Millisecond), source(10, 15, 3*time.Millisecond), source(15, 40, 0)))
	if len(got) != 40 {
		t.Fatalf("received %d values, want 40", len(got))
	}
	for i, v := range got {
		if v != i {
			t.Fatalf("value %d missing or duplicated (got %v)", i, got)
		}
	}
}

func TestMergeNothing(t *testing.T) {
	if got := drain(t, Merge()); len(got) != 0 {
		t.Fatalf("Merge() delivered %v", got)
	}
}
