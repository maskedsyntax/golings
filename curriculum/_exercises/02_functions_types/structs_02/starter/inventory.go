// structs_02: A slice of structs
//
// Implement Total, Cheapest and ByName. See README.md.
//
//	items := []Item{{"pen", 2, 10}, {"book", 15, 1}}
//	Total(items) == 35
//	Cheapest(items) == Item{"pen", 2, 10}, true
//
// Reference: https://gobyexample.com/structs
package inventory

// Item is a line in an inventory.
type Item struct {
	Name  string
	Price int // in cents
	Qty   int
}

// Total returns the combined value of all items.
func Total(items []Item) int {
	// TODO
	return 0
}

// Cheapest returns the first item with the lowest price. ok is false if items is empty.
func Cheapest(items []Item) (Item, bool) {
	// TODO
	return Item{}, false
}

// ByName indexes items by name.
func ByName(items []Item) map[string]Item {
	// TODO
	return nil
}
