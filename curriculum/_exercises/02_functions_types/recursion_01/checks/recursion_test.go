package recursion

import "testing"

func TestFactorial(t *testing.T) {
	for _, tt := range []struct{ n, want int }{{0, 1}, {1, 1}, {2, 2}, {5, 120}, {10, 3628800}} {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFib(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for n, w := range want {
		if got := Fib(n); got != w {
			t.Errorf("Fib(%d) = %d, want %d", n, got, w)
		}
	}
	if got := Fib(20); got != 6765 {
		t.Errorf("Fib(20) = %d, want 6765", got)
	}
}
