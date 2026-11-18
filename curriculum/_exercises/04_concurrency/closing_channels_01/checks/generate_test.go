package generate

import (
	"slices"
	"testing"
	"time"
)

func drain(t *testing.T, ch <-chan int) []int {
	t.Helper()
	var out []int
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return out
			}
			out = append(out, v)
		case <-timeout:
			t.Fatalf("received %v and then waited 2s: the channel was never closed", out)
		}
	}
}

func TestGenerateCloses(t *testing.T) {
	if got := drain(t, Generate(1, 2, 3)); !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("received %v, want [1 2 3]", got)
	}
}

func TestGenerateNothing(t *testing.T) {
	if got := drain(t, Generate()); len(got) != 0 {
		t.Fatalf("received %v, want nothing", got)
	}
}
