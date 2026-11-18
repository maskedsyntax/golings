# generics_01: Constraints

A type parameter's **constraint** says which types it may be, and therefore which operations
the function body may use. `any` allows every type, so you can't add two `any` values.

Constraints are interfaces, and they may list types:

```go
type Number interface {
	~int | ~int64 | ~float64 // ~int: int, or any type defined as int
}
```

## Your task

`Sum` doesn't compile. Declare a `Number` constraint and use it, so that `Sum` works on
`[]int`, `[]int64`, `[]float64`, and slices of named types such as `Cents` (defined as `int`).

Reference: https://gobyexample.com/generics
