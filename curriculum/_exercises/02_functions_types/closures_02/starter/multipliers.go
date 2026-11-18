// closures_02: Closures capture variables
//
// Every function returned by Multipliers multiplies by the last factor. Fix it.
//
//	fns := Multipliers([]int{2, 3})
//	fns[0](5) == 10; fns[1](5) == 15
//
// Reference: https://gobyexample.com/closures
package multipliers

// Multipliers returns one function per factor; the i-th multiplies by factors[i].
func Multipliers(factors []int) []func(int) int {
	var fns []func(int) int
	var factor int
	for _, f := range factors {
		factor = f
		fns = append(fns, func(x int) int { return x * factor })
	}
	return fns
}
