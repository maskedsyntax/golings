// range_over_iterators_01: Writing an iterator (reference solution)
package countdown

import "iter"

// Countdown yields n, n-1, ..., 1.
func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i >= 1; i-- {
			if !yield(i) {
				return // the caller's loop stopped
			}
		}
	}
}
