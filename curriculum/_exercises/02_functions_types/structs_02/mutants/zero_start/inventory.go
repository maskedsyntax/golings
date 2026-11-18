// Mutant: starts Cheapest from the zero Item.
package inventory

type Item struct {
	Name  string
	Price int
	Qty   int
}

func Total(items []Item) int {
	total := 0
	for _, it := range items {
		total += it.Price * it.Qty
	}
	return total
}

func Cheapest(items []Item) (Item, bool) {
	var best Item
	for _, it := range items {
		if best.Name == "" || it.Price < best.Price {
			best = it
		}
	}
	return best, len(items) > 0
}

func ByName(items []Item) map[string]Item {
	var m map[string]Item
	for _, it := range items {
		if m == nil {
			m = map[string]Item{}
		}
		m[it.Name] = it
	}
	return m
}
