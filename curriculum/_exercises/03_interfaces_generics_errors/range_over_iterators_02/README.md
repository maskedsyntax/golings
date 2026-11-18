# range_over_iterators_02: Composing iterators

Iterators are lazy: values are produced only when the consumer asks for the next one. That
means you can build pipelines over infinite sequences:

```go
for n := range Take(Filter(Naturals(), isOdd), 3) {
	fmt.Println(n) // 1, 3, 5
}
```

## Your task

`Naturals()` yields `0, 1, 2, …` forever. Implement two generic adapters:

- `Filter(seq, keep)` yields only the values for which `keep` returns true.
- `Take(seq, n)` yields at most the first `n` values of `seq`, then stops. For `n <= 0` it
  yields nothing, and pulls nothing from `seq`.

Both must stop as soon as their consumer stops.

Reference: https://gobyexample.com/range-over-iterators
