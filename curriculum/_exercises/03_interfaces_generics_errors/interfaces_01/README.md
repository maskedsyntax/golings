# interfaces_01: Satisfying an interface

An interface is a set of method signatures. Any type with those methods satisfies it,
**implicitly**: there is no `implements` declaration.

```go
type geometry interface {
	area() float64
	perim() float64
}
```

## Your task

`TotalPerimeter` accepts a `[]Shape`, and the tests pass it both rectangles and circles.
The package doesn't compile, because `Circle` doesn't satisfy `Shape` yet. Fix that.

Reference: https://gobyexample.com/interfaces
