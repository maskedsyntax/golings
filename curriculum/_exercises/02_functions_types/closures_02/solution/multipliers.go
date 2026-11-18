// closures_02: Closures capture variables (reference solution)
package multipliers

// Multipliers returns one function per factor; the i-th multiplies by factors[i].
func Multipliers(factors []int) []func(int) int {
	fns := make([]func(int) int, 0, len(factors))
	for _, f := range factors {
		// Since Go 1.22 each iteration has its own f, so each closure
		// captures a different variable.
		fns = append(fns, func(x int) int { return x * f })
	}
	return fns
}
