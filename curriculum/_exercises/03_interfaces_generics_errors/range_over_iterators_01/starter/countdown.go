// range_over_iterators_01: Writing an iterator
//
// Implement Countdown. See README.md.
//
//	slices.Collect(Countdown(3)) == []int{3, 2, 1}
//
// Reference: https://gobyexample.com/range-over-iterators
package countdown

import "iter"

// Countdown yields n, n-1, ..., 1.
func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// TODO
	}
}
