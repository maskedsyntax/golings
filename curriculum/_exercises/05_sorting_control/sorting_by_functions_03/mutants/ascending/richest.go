// Mutant: sorts smallest first.
package richest

import (
	"cmp"
	"slices"
)

type Account struct {
	Owner   string
	Balance int64
}

func Richest(accts []Account) {
	slices.SortFunc(accts, func(a, b Account) int { return cmp.Compare(a.Balance, b.Balance) })
}
