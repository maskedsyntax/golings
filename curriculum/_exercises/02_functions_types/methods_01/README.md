# methods_01: Receivers

A **method** is a function with a receiver:

```go
func (r rect) area() int     { … } // value receiver: works on a copy
func (r *rect) grow(n int)   { … } // pointer receiver: can modify the caller's rect
```

## Your task

The tests call `r.Area()` and `r.Scale(f)`, but `rect.go` declares plain functions. Turn them
into methods:

- `Area` returns `Width × Height`.
- `Scale(f)` multiplies both sides by `f`, **changing the receiver**.

Reference: https://gobyexample.com/methods
