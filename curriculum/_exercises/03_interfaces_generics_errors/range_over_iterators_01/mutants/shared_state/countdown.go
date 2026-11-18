// Mutant: the counter lives outside the iterator, so a second range yields nothing.
package countdown

import "iter"

func Countdown(n int) iter.Seq[int] {
	i := n
	return func(yield func(int) bool) {
		for ; i >= 1; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
