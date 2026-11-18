package square

import (
	"testing"
	"time"
)

func feed(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func recv(t *testing.T, ch <-chan int) int {
	t.Helper()
	if ch == nil {
		t.Fatal("Square returned a nil channel")
	}
	select {
	case v := <-ch:
		return v
	case <-time.After(2 * time.Second):
		t.Fatal("no value arrived within 2s")
		return 0
	}
}

func TestSquare(t *testing.T) {
	out := Square(feed(1, 2, 3, -4))
	for _, want := range []int{1, 4, 9, 16} {
		if got := recv(t, out); got != want {
			t.Fatalf("received %d, want %d", got, want)
		}
	}
}

func TestSquareChains(t *testing.T) {
	out := Square(Square(feed(2, 3)))
	for _, want := range []int{16, 81} {
		if got := recv(t, out); got != want {
			t.Fatalf("Square(Square(...)) delivered %d, want %d", got, want)
		}
	}
}
