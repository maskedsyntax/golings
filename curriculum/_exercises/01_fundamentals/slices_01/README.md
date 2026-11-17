# slices_01: Reverse without mutating

Slices share their backing array, so changing elements through one slice changes them for
everyone who holds that slice.

## Your task

Implement `Reverse(in []int) []int`. It must:

- return a **new** slice with the elements in reverse order,
- **never modify** `in`,
- return an empty, **non-nil** slice for empty or nil input.

Reference: https://gobyexample.com/slices
