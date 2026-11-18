# arrays_02: Matrices as arrays

`[3][3]int` is an array of three `[3]int` arrays: a fixed 3×3 matrix. Arrays whose elements
are comparable can be compared with `==`.

## Your task

Implement in `matrix.go`:

- `Transpose(m [3][3]int) [3][3]int` returns `m` with rows and columns swapped.
- `Equal(a, b [3][3]int) bool` reports whether two matrices are identical.

Reference: https://gobyexample.com/arrays
