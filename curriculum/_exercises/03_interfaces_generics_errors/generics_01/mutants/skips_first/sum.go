// Mutant: starts summing from the second element.
package sum

type Cents int

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](xs []T) T {
	var total T
	for i := 1; i < len(xs); i++ {
		total += xs[i]
	}
	return total
}
