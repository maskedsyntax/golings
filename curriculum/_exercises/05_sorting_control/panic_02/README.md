# panic_02: Panics you didn't ask for

Some panics come from the runtime itself, when code does something impossible: indexing past
the end of a slice, writing to a nil map, dereferencing a nil pointer.

Both functions in this file panic on perfectly reasonable input.

## Your task

Fix them so they never panic:

- `Last` returns `0, false` for an empty slice.
- A zero-value `Registry` (`var r Registry`) must be ready to use.

Reference: https://gobyexample.com/panic
