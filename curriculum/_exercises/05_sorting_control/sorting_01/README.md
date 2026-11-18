# sorting_01: A sorted copy

The `slices` package sorts any slice of ordered values (numbers, strings) **in place**:

```go
strs := []string{"c", "a", "b"}
slices.Sort(strs)
fmt.Println(strs) // [a b c]
```

`Alphabetize` should return the words in alphabetical order **without changing the caller's
slice**. It does return sorted words, but it also scrambles the caller's slice.

## Your task

Fix `Alphabetize` so the input is left as it was.

Reference: https://gobyexample.com/sorting
