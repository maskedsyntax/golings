// recursion_01: Base case and recursive case (reference solution)
package recursion

// Factorial returns n! for n >= 0.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fib returns the n-th Fibonacci number, with Fib(0) == 0 and Fib(1) == 1.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
