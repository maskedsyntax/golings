# maps_03: Grouping into a map of slices

A missing key gives you the zero value, and the zero value of a slice is `nil`, which `append`
accepts. That makes "group by" a one-liner per item.

## Your task

Implement `GroupByLength(words []string) map[int][]string`. It groups words by their length in
**characters** (runes, so `"héllo"` has length 5), keeping the original order within each
group. It always returns a non-nil map.

```text
GroupByLength([go is fun ok]) == map[2:[go is ok] 3:[fun]]
```

Reference: https://gobyexample.com/maps
