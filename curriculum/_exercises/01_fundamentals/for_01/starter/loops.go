// for_01: Go has only for
//
// This file is written with habits from other languages and does not
// compile. Rewrite the loops the Go way.
//
// Reference: https://gobyexample.com/for
package loops

// CountDown returns n, n-1, ..., 1.
func CountDown(n int) []int {
	var out []int
	while n > 0 {
		out = append(out, n)
		n--
	}
	return out
}

// SumTo returns 1 + 2 + ... + n.
func SumTo(n int) int {
	total := 0
	for (i := 1; i <= n; i++) {
		total += i
	}
	return total
}
