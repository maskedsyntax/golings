package pool

import (
	"slices"
	"sync/atomic"
	"testing"
	"time"
)

func TestMapPool(t *testing.T) {
	var active, peak atomic.Int32
	f := func(n int) int {
		cur := active.Add(1)
		for {
			p := peak.Load()
			if cur <= p || peak.CompareAndSwap(p, cur) {
				break
			}
		}
		time.Sleep(5 * time.Millisecond)
		active.Add(-1)
		return n * 10
	}
	items := make([]int, 40)
	want := make([]int, 40)
	for i := range items {
		items[i] = i
		want[i] = i * 10
	}
	got := MapPool(items, 4, f)
	if !slices.Equal(got, want) {
		t.Fatalf("MapPool returned %v, want %v", got, want)
	}
	switch p := peak.Load(); {
	case p > 4:
		t.Fatalf("%d calls ran at once; the limit is 4 workers", p)
	case p < 2:
		t.Fatalf("only %d call ran at a time; use all 4 workers", p)
	}
}

func TestMapPoolMoreWorkersThanItems(t *testing.T) {
	if got := MapPool([]int{1, 2}, 8, func(n int) int { return -n }); !slices.Equal(got, []int{-1, -2}) {
		t.Fatalf("MapPool = %v, want [-1 -2]", got)
	}
}
