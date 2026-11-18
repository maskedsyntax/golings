# pointers_02: nil pointers and identity

A pointer's zero value is `nil`, and dereferencing `nil` panics. Code that accepts pointers
usually checks for `nil` first.

Two pointers compare equal only when they point to the **same variable**. Pointing to two
variables that happen to hold equal values isn't enough.

## Your task

Implement, in `pointers.go`:

- `ValueOr(p *int, def int) int` returns `*p`, or `def` if `p` is nil.
- `Increment(p *int)` adds 1 to `*p`, and does nothing if `p` is nil.
- `Largest(ptrs []*int) *int` returns the pointer (from `ptrs`) to the largest value. It skips
  nil entries, and returns nil if there's nothing to compare. If several values tie, return the
  first one.

Reference: https://gobyexample.com/pointers
