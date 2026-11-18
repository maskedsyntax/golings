# closures_02: Closures capture variables

A closure doesn't copy the variables it uses. It refers to them. If the variable changes
later, the closure sees the new value.

## Your task

`Multipliers(factors []int) []func(int) int` should return one function per factor, where
the i-th function multiplies its argument by `factors[i]`:

```go
fns := Multipliers([]int{2, 3, 10})
fns[0](5) // 10
fns[1](5) // 15
```

Every returned function currently behaves like the last one. Find out why and fix it.

Reference: https://gobyexample.com/closures
