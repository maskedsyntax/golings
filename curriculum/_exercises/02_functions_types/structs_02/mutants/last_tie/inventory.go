// Mutant: <= makes the last of several equally cheap items win.
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
	if len(items) == 0 {
		return Item{}, false
	}
	best := items[0]
	for _, it := range items[1:] {
		if it.Price <= best.Price {
			best = it
		}
	}
	return best, true
}

func ByName(items []Item) map[string]Item {
	m := make(map[string]Item, len(items))
	for _, it := range items {
		m[it.Name] = it
	}
	return m
}
