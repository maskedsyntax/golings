// Mutant: the closures read factors[i] when called, so later changes leak in.
package multipliers

func Multipliers(factors []int) []func(int) int {
	var fns []func(int) int
	for i := range factors {
		fns = append(fns, func(x int) int { return x * factors[i] })
	}
	return fns
}
