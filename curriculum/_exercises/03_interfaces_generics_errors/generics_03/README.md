# generics_03: A generic set

Types can have type parameters too:

```go
type List[T any] struct { … }
func (l *List[T]) Push(v T) { … }
```

Map keys must support `==`, so a set's element type is constrained by the built-in
`comparable`.

## Your task

Implement `Set[T comparable]`:

- `Add(v)` adds `v`, and adding it again has no effect.
- `Has(v)` reports whether `v` is in the set.
- `Len()` returns the number of distinct elements.
- `Union(a, b)` returns a **new** set with the elements of both. `a` and `b` are unchanged.

`var s Set[string]` must be usable straight away, with no constructor.

Reference: https://gobyexample.com/generics
