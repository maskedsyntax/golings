package pipe

import (
	"slices"
	"testing"
	"time"
)

func within(t *testing.T, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s blocked for 2s", what)
	}
}

func TestGenerate(t *testing.T) {
	var ch <-chan int
	within(t, "Generate", func() { ch = Generate(1, 2, 3) })
	if ch == nil {
		t.Fatal("Generate returned a nil channel")
	}
	for want := 1; want <= 3; want++ {
		select {
		case got := <-ch:
			if got != want {
				t.Fatalf("received %d, want %d", got, want)
			}
		case <-time.After(2 * time.Second):
			t.Fatalf("value %d was never sent", want)
		}
	}
}

func TestTake(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 10; ; i++ {
			ch <- i
		}
	}()
	var got []int
	within(t, "Take", func() { got = Take(ch, 3) })
	if !slices.Equal(got, []int{10, 11, 12}) {
		t.Fatalf("Take(ch, 3) = %v, want [10 11 12]", got)
	}
}

func TestTogether(t *testing.T) {
	var got []int
	within(t, "Take(Generate(...))", func() { got = Take(Generate(4, 5, 6), 2) })
	if !slices.Equal(got, []int{4, 5}) {
		t.Fatalf("Take(Generate(4, 5, 6), 2) = %v, want [4 5]", got)
	}
}
