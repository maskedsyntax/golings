# switch_03: Type switches

A value of type `any` can hold anything. A **type switch** branches on what it actually holds:

```go
switch x := v.(type) {
case int:    // x is an int here
case string: // x is a string here
}
```

## Your task

Implement `Describe(v any) string`:

| v | result |
|---|---|
| `nil` | `"nil"` |
| `42` | `"int 42"` |
| `"hi"` | `` `string "hi"` `` (quoted) |
| `true` | `"bool true"` |
| `[]int{1, 2, 3}` | `"[]int of length 3"` |
| anything else | `"unknown"` |

Reference: https://gobyexample.com/switch
