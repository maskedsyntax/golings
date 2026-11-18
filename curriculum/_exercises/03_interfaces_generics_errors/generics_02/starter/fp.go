// generics_02: Map, Filter, Reduce
//
// Implement the three generic functions. See README.md.
//
//	Map([]string{"a", "bb"}, func(s string) int { return len(s) }) == []int{1, 2}
//	Reduce([]int{1, 2, 3}, 10, func(acc, x int) int { return acc + x }) == 16
//
// Reference: https://gobyexample.com/generics
package fp

// Map returns f(x) for every x in xs.
func Map[T, U any](xs []T, f func(T) U) []U {
	// TODO
	return nil
}

// Filter returns the elements of xs for which keep is true.
func Filter[T any](xs []T, keep func(T) bool) []T {
	// TODO
	return nil
}

// Reduce folds xs from the left, starting from init.
func Reduce[T, A any](xs []T, init A, f func(A, T) A) A {
	// TODO
	return init
}
