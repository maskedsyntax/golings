# values_01: Mixing numbers and strings

Go has strings, integers, floats and booleans, and it **never converts between
them implicitly**. This file doesn't compile because it mixes types.

## Your task

Fix `values.go` so it compiles and behaves as documented:

- `Average(sum, count int) float64` returns the exact average, so `Average(7, 2)` is `3.5`.
  A `count` of 0 returns 0.
- `Label(name string, age int) string` returns text like `"Gopher is 13"`.

Reference: https://gobyexample.com/values
