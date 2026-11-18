# structs_02: A slice of structs

Structs are values: assigning one or passing it to a function copies all its fields.
Slices and maps of structs are the everyday way to work with records.

## Your task

`Item` is defined in `inventory.go`. Implement:

- `Total(items []Item) int` returns the sum of `Price × Qty` over all items.
- `Cheapest(items []Item) (Item, bool)` returns the item with the lowest `Price`, and `true`.
  For a tie, it returns the first such item. For an empty slice, it returns `Item{}, false`.
- `ByName(items []Item) map[string]Item` indexes the items by `Name`. If two items have the
  same name, the later one wins.

Reference: https://gobyexample.com/structs
