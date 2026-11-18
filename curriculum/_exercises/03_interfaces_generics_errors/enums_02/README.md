# enums_02: Names out of sync

Enum-style constants often come with a table of names. When the table is a plain slice,
inserting a new constant in the middle of the `const` block quietly shifts every later
name.

## Your task

`Suspended` was added recently. Since then, `Suspended.String()` returns `"closed"` and
`ParseStatus("closed")` returns the wrong value. Fix `status.go` so that:

- `String()` returns the lowercase name of each constant, or `"Status(n)"` for unknown values,
- `ParseStatus(name)` is the exact inverse of `String()` for the four statuses, and returns
  `false` for anything else.

Make it hard for the next new constant to break things again.

Reference: https://gobyexample.com/enums
