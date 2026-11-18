# generics_04: Generic types and cmp.Ordered

A struct can be generic, like `Pair[K, V]` below. The standard library provides useful
constraints: `comparable` for map keys, and
[`cmp.Ordered`](https://pkg.go.dev/cmp#Ordered) for anything you can sort with `<`.

## Your task

Implement, in `pairs.go`:

- `Entries(m map[K]V) []Pair[K, V]` returns the map's entries **sorted by key**.
- `GroupBy(xs []T, key func(T) K) map[K][]T` groups the elements by `key(x)`, keeping the
  original order inside each group.

```go
Entries(map[string]int{"b": 2, "a": 1})   // [{a 1} {b 2}]
GroupBy([]string{"go", "c", "js"}, length) // {1: [c], 2: [go js]}
```

Reference: https://gobyexample.com/generics
