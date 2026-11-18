// Mutant: keeps calling yield after it returned false.
package countdown

import "iter"

func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i >= 1; i-- {
			yield(i)
		}
	}
}
