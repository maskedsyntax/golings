# interfaces_02: Type switches

A value of interface type (such as `any`) carries a concrete type at run time. A **type
switch** branches on it:

```go
switch v := x.(type) {
case int:
	// v is an int here
case fmt.Stringer:
	// v is a fmt.Stringer here
}
```

## Your task

Implement `Describe(x any) string` with a type switch:

| `x` holds | Result |
|---|---|
| nothing (`nil`) | `"nil"` |
| an `int` | `"int 42"` |
| a `string` | `` `string "hi"` `` (quoted, as `%q` does) |
| a `[]int` | `"3 ints"` (its length) |
| any other `fmt.Stringer` | `"stringer "` + its `String()` |
| anything else | `"unknown"` |

Reference: https://gobyexample.com/interfaces
