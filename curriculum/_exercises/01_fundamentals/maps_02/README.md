# maps_02: Inventory, `delete` and sorted keys

A map passed to a function is shared with the caller, not copied. `delete(m, k)` removes a key,
and ranging over a map visits keys in an **unspecified, deliberately varying order**.

## Your task

Implement in `inventory.go`, where `stock` maps item names to quantities:

- `Sell(stock map[string]int, item string, n int) bool` removes `n` units if at least `n` are
  available, and returns whether the sale happened. When an item's quantity reaches 0, its key
  is **removed** from the map.
- `Items(stock map[string]int) []string` returns the item names **in alphabetical order**. It
  returns an empty, non-nil slice for an empty map.

Reference: https://gobyexample.com/maps
