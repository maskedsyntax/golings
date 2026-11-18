// variables_02: The shadowed balance
//
// Balance compiles but is wrong. Find out why the named results never change.
//
//	Balance([]int{10, -20, 30}) == (20, true)
//
// Reference: https://gobyexample.com/variables
package variables

// Balance sums transactions and reports whether the running balance ever
// went below zero.
func Balance(transactions []int) (balance int, overdrawn bool) {
	for _, t := range transactions {
		balance := balance + t
		if balance < 0 {
			overdrawn := true
			_ = overdrawn
		}
	}
	return balance, overdrawn
}
