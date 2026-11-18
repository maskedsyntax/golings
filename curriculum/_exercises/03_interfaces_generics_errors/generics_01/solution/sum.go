// generics_01: Constraints (reference solution)
package sum

// Cents is an amount of money.
type Cents int

// Number is the set of types Sum can add.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum returns the sum of xs.
func Sum[T Number](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}
