// structs_02: A slice of structs (reference solution)
package inventory

// Item is a line in an inventory.
type Item struct {
	Name  string
	Price int // in cents
	Qty   int
}

// Total returns the combined value of all items.
func Total(items []Item) int {
	total := 0
	for _, it := range items {
		total += it.Price * it.Qty
	}
	return total
}

// Cheapest returns the first item with the lowest price. ok is false if items is empty.
func Cheapest(items []Item) (Item, bool) {
	if len(items) == 0 {
		return Item{}, false
	}
	best := items[0]
	for _, it := range items[1:] {
		if it.Price < best.Price {
			best = it
		}
	}
	return best, true
}

// ByName indexes items by name.
func ByName(items []Item) map[string]Item {
	m := make(map[string]Item, len(items))
	for _, it := range items {
		m[it.Name] = it
	}
	return m
}
