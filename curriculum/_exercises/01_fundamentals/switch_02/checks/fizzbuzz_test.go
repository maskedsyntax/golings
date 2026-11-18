package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	for _, tt := range []struct {
		n    int
		want string
	}{
		{1, "1"}, {3, "Fizz"}, {5, "Buzz"}, {9, "Fizz"}, {10, "Buzz"}, {15, "FizzBuzz"},
		{30, "FizzBuzz"}, {7, "7"}, {0, "FizzBuzz"}, {-3, "Fizz"},
	} {
		if got := FizzBuzz(tt.n); got != tt.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
