# values_02: Booleans and operators

Go's `bool` values combine with `&&` (and), `||` (or) and `!` (not). Comparisons such as `==`
and `>=` produce booleans.

## Your task

Implement in `values.go`:

- `CanVote(age int, registered bool) bool` is true for registered people aged 18 or over.
- `Differ(a, b bool) bool` is true when exactly one of `a` and `b` is true (exclusive or).
- `IsEven(n int) bool` reports whether `n` is even. It must work for negative numbers too.

Reference: https://gobyexample.com/values
