# functions_01: Parameter and result types

A Go function declares the type of each parameter and the type of what it returns:

```go
func plus(a int, b int) int { return a + b }
```

Go has no implicit returns: a function with a result must `return` it explicitly.

## Your task

`functions.go` doesn't compile. Fix the two function declarations so that:

- `Add(2, 3)` returns `5`,
- `Greeting("Gopher")` returns `"Hello, Gopher!"`.

Reference: https://gobyexample.com/functions
