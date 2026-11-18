# range_over_iterators_03: Iterating a custom collection

Iterators let your own data structures work with `range`, just like slices and maps. An
[`iter.Seq2[K, V]`](https://pkg.go.dev/iter#Seq2) yields pairs:

```go
for i, v := range list.All() { … }
```

The standard library follows a naming convention: `All` yields everything, `Values` only
values, and `Backward` iterates in reverse.

## Your task

`List[T]` is a singly linked list with `Push` and `Values` already written. Implement:

- `All() iter.Seq2[int, T]` yields `(index, value)` pairs from front to back.
- `Backward() iter.Seq2[int, T]` yields the same pairs from back to front:
  `(len-1, last) … (0, first)`.

Both must stop when the loop body breaks.

Reference: https://gobyexample.com/range-over-iterators
