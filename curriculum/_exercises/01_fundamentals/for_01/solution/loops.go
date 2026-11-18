// for_01: Go has only for (reference solution)
package loops

// CountDown returns n, n-1, ..., 1.
func CountDown(n int) []int {
	var out []int
	for n > 0 {
		out = append(out, n)
		n--
	}
	return out
}

// SumTo returns 1 + 2 + ... + n.
func SumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}
