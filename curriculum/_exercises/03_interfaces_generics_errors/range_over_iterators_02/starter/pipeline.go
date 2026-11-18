// range_over_iterators_02: Composing iterators
//
// Implement Filter and Take. See README.md.
//
//	slices.Collect(Take(Filter(Naturals(), isOdd), 3)) == []int{1, 3, 5}
//
// Reference: https://gobyexample.com/range-over-iterators
package pipeline

import "iter"

// Naturals yields 0, 1, 2, ... forever.
func Naturals() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Filter yields the values of seq for which keep returns true.
func Filter[T any](seq iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		// TODO
	}
}

// Take yields at most the first n values of seq.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		// TODO
	}
}
