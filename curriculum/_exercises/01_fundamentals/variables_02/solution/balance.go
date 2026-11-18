// variables_02: The shadowed balance (reference solution)
package variables

// Balance sums transactions and reports whether the running balance ever
// went below zero.
func Balance(transactions []int) (balance int, overdrawn bool) {
	for _, t := range transactions {
		balance += t
		if balance < 0 {
			overdrawn = true
		}
	}
	return balance, overdrawn
}
