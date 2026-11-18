# sorting_by_functions_01: Shortest first

To sort by something other than the natural order, give `slices.SortFunc` a comparison
function:

```go
lenCmp := func(a, b string) int {
	return cmp.Compare(len(a), len(b)) // <0: a first, >0: b first, 0: equal
}
slices.SortFunc(fruits, lenCmp)
```

`ByLength` should sort words from shortest to longest, with words of the same length in
alphabetical order. It sorts alphabetically instead.

## Your task

Fix `ByLength`. `cmp.Or` returns its first non-zero argument, which is handy for "then by …"
tie-breaks.

Reference: https://gobyexample.com/sorting-by-functions
