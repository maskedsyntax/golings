// generics_01: Constraints
//
// This doesn't compile: + isn't defined for every type T. Fix the constraint.
//
//	Sum([]int{1, 2, 3}) == 6
//	Sum([]float64{0.5, 0.25}) == 0.75
//	Sum([]Cents{150, 50}) == Cents(200)
//
// Reference: https://gobyexample.com/generics
package sum

// Cents is an amount of money.
type Cents int

// Sum returns the sum of xs.
func Sum[T any](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}
