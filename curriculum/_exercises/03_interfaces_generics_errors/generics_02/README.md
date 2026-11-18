# generics_02: Map, Filter, Reduce

A function can have several type parameters:

```go
func Map[T, U any](xs []T, f func(T) U) []U
```

The compiler usually **infers** them from the arguments, so callers just write
`Map(words, strings.ToUpper)`.

## Your task

Implement in `fp.go`:

- `Map(xs, f)` returns a new slice of `f(x)` for each `x`.
- `Filter(xs, keep)` returns a new slice of the `x` for which `keep(x)` is true.
- `Reduce(xs, init, f)` folds the slice from left to right: `f(…f(f(init, xs[0]), xs[1])…)`.
  For an empty slice it returns `init`.

None of them may modify `xs`.

Reference: https://gobyexample.com/generics
