# range_over_built_in_types_02: Ranging over ints and maps

Since Go 1.22, `for i := range 5` loops with `i = 0, 1, 2, 3, 4`. Ranging over a map visits
every key exactly once, but **in no particular order**. Go randomizes it on purpose, so code
can't come to depend on it.

## Your task

Implement, in `ranges.go`:

- `Squares(n int) []int` returns `[1, 4, 9, …, n²]` (empty for `n <= 0`). Use `range` over
  the integer `n`.
- `SortedKeys(m map[string]int) []string` returns the keys of `m` in alphabetical order.

This exercise checks that `Squares` uses a range loop.

Reference: https://gobyexample.com/range-over-built-in-types
