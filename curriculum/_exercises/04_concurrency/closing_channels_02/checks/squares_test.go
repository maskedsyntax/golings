package squares

import (
	"slices"
	"testing"
	"time"
)

func collect(t *testing.T, ch <-chan int) []int {
	t.Helper()
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
			t.Fatalf("results was never closed (got %d values)", len(out))
		}
	}
}

func TestSquareAll(t *testing.T) {
	nums := make([]int, 100)
	want := make([]int, 100)
	for i := range nums {
		nums[i] = i
		want[i] = i * i
	}
	for _, workers := range []int{1, 4, 10} {
		if got := collect(t, SquareAll(nums, workers)); !slices.Equal(got, want) {
			t.Fatalf("with %d workers: got %d values %v..., want the 100 squares", workers, len(got), got[:min(5, len(got))])
		}
	}
}
