# if_else_02: `if` with a short statement

An `if` can start with a statement that runs before the condition:

```go
if n, err := strconv.Atoi(s); err != nil {
    // n and err are in scope in every branch...
}
// ...but not here
```

## Your task

Implement `Classify(s string) string`. It parses `s` as an integer (ignoring surrounding
spaces) and returns `"not a number"`, `"negative"`, `"zero"` or `"positive"`.

Reference: https://gobyexample.com/if-else
