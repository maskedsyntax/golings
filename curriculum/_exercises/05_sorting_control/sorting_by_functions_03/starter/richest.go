// sorting_by_functions_03: Subtraction is not comparison
//
// Richest gets the order wrong for large balances. Fix it. See README.md.
//
// Reference: https://gobyexample.com/sorting-by-functions
package richest

import "slices"

// Account is a balance in cents.
type Account struct {
	Owner   string
	Balance int64
}

// Richest sorts accounts in place, largest balance first.
func Richest(accts []Account) {
	slices.SortFunc(accts, func(a, b Account) int {
		return int(b.Balance - a.Balance)
	})
}
