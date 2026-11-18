// Mutant: sold-out items stay in the map with quantity 0.
package inventory

import "slices"

func Sell(stock map[string]int, item string, n int) bool {
	if n <= 0 || stock[item] < n {
		return false
	}
	stock[item] -= n
	return true
}

func Items(stock map[string]int) []string {
	keys := make([]string, 0, len(stock))
	for k := range stock {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
