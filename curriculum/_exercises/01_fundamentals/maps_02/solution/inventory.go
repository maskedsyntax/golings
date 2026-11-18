// maps_02: Inventory: delete and sorted keys (reference solution)
package inventory

import "slices"

// Sell removes n units of item if enough are in stock, deleting the item
// when none remain. It reports whether the sale happened.
func Sell(stock map[string]int, item string, n int) bool {
	if n <= 0 || stock[item] < n {
		return false
	}
	stock[item] -= n
	if stock[item] == 0 {
		delete(stock, item)
	}
	return true
}

// Items returns the item names in alphabetical order.
func Items(stock map[string]int) []string {
	keys := make([]string, 0, len(stock))
	for k := range stock {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
