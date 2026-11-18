package memoize

import "testing"

func TestMemoizeCallsOncePerArgument(t *testing.T) {
	calls := map[int]int{}
	square := Memoize(func(n int) int { calls[n]++; return n * n })
	for _, n := range []int{4, 4, 3, 4, 0, 0, 3, -2, -2} {
		if got := square(n); got != n*n {
			t.Errorf("square(%d) = %d, want %d", n, got, n*n)
		}
	}
	for n, c := range calls {
		if c != 1 {
			t.Errorf("f(%d) was called %d times, want 1", n, c)
		}
	}
}

func TestMemoizedFunctionsHaveSeparateCaches(t *testing.T) {
	double := Memoize(func(n int) int { return n * 2 })
	negate := Memoize(func(n int) int { return -n })
	if got := double(5); got != 10 {
		t.Errorf("double(5) = %d, want 10", got)
	}
	if got := negate(5); got != -5 {
		t.Errorf("negate(5) = %d, want -5 (is the cache shared between memoized functions?)", got)
	}
}

func TestCachedZero(t *testing.T) {
	calls := 0
	zero := Memoize(func(int) int { calls++; return 0 })
	zero(1)
	zero(1)
	if calls != 1 {
		t.Errorf("a result of 0 was recomputed: f called %d times, want 1", calls)
	}
}
