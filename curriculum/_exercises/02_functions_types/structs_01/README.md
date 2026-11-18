# structs_01: Fields and struct literals

A struct groups named fields:

```go
type person struct {
	name string
	age  int
}
p := person{name: "Ada", age: 36}
pp := &person{name: "Alan"} // a pointer to a new struct
```

## Your task

`point.go` doesn't compile. Fix it so that:

- `Point` has two `int` fields, `X` and `Y`,
- `NewPoint(x, y)` returns a pointer to a new `Point`,
- `Manhattan(a, b)` (already written) works.

Reference: https://gobyexample.com/structs
