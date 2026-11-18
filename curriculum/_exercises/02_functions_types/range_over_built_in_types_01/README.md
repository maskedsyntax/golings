# range_over_built_in_types_01: What range gives you

`range` yields different things for different types:

| Ranging over | First value | Second value |
|---|---|---|
| slice or array | index | element |
| string | byte offset | **rune** (decoded character) |
| map | key | value |

## Your task

Both functions in `ranges.go` compile but return wrong answers:

- `Sum(nums []int) int` should add up the elements.
- `CountRunes(s string) map[rune]int` should count each character, including non-ASCII ones:
  `CountRunes("héé")` → `{'h': 1, 'é': 2}`.

Fix them.

Reference: https://gobyexample.com/range-over-built-in-types
