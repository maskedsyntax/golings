# functions_02: Functions as values

Functions are values in Go. They can be passed to other functions, just like ints or
strings. The parameter's type spells out the function's signature: `f func(int) int`.

## Your task

Implement, in `apply.go`:

- `Map(nums []int, f func(int) int) []int` returns a **new** slice with `f` applied to each
  element, in order.
- `Filter(nums []int, keep func(int) bool) []int` returns a new slice with only the elements
  for which `keep` returns true, in order.

Neither function may modify `nums`. For an empty input, both return an empty (or nil) slice.

Reference: https://gobyexample.com/functions
