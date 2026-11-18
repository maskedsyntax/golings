# range_over_iterators_01: Writing an iterator

Since Go 1.23, `range` works over functions. An
[`iter.Seq[V]`](https://pkg.go.dev/iter) is a function that takes a `yield` callback and
calls it once per value:

```go
func(yield func(V) bool)
```

Each call to `yield` runs the body of the caller's `for … range` loop. If that body
executes `break` (or `return`), `yield` returns `false`, and the iterator **must stop**.

## Your task

Implement `Countdown(n int) iter.Seq[int]`, which yields `n, n-1, …, 1`. It yields nothing
for `n <= 0`.

```go
for i := range Countdown(3) {
	fmt.Println(i) // 3, 2, 1
}
```

It must stop as soon as the caller's loop stops.

Reference: https://gobyexample.com/range-over-iterators
