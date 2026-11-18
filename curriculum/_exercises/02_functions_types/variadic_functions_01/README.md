# variadic_functions_01: Spreading a slice

A **variadic** function accepts any number of trailing arguments of one type:

```go
func sum(nums ...int) int { … }
sum(1, 2, 3)
```

Inside the function, `nums` is a `[]int`. If you already have a slice, pass it with
`sum(xs...)`.

## Your task

`sum.go` doesn't compile. `Sum` is fine; fix the two functions that call it so that:

- `Average([]int{2, 4, 9})` returns `5`, and `Average(nil)` returns `0`,
- `Total(1, 2, 3)` returns `6`.

Reference: https://gobyexample.com/variadic-functions
