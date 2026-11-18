# methods_02: A stack with a useful zero value

Idiomatic Go types often work with no constructor: their **zero value** is ready to use.
A struct wrapping a nil slice gets this almost for free, because `append` and `len` both
handle nil slices.

## Your task

Implement the methods of `Stack`, a last-in, first-out stack of ints:

- `Push(v int)` adds `v` on top.
- `Pop() (int, bool)` removes and returns the top value. On an empty stack, it returns `0, false`.
- `Peek() (int, bool)` returns the top value without removing it.
- `Len() int` returns the number of values.

`var s Stack` must be usable straight away. Think about which methods need pointer receivers.

Reference: https://gobyexample.com/methods
