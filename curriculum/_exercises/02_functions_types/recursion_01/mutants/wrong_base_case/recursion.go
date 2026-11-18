// Mutant: Fib(0) returns 1, shifting the whole sequence.
package recursion

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fib(n int) int {
	if n < 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
