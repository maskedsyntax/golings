# maps_01: Counting and membership

A map lookup for a missing key returns the zero value, so `m["nope"]` is `0` for a
`map[string]int`. To tell "present with value 0" apart from "absent", use the two-value
"comma ok" form.

## Your task

Implement in `maps.go`:

- `WordCount(words []string) map[string]int` counts how often each word appears. It's
  case-sensitive, and it always returns a non-nil map, even for no words.
- `Has(m map[string]int, key string) bool` reports whether `key` is **present**, even if its
  value is 0.

Reference: https://gobyexample.com/maps
