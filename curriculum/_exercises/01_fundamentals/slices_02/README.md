# slices_02: Chunks that don't overlap

`s[low:high]` creates a view of `s` whose capacity extends to the end of `s`. Appending to
that view can quietly overwrite the elements that follow it. The three-index form
`s[low:high:max]` limits the capacity.

## Your task

Implement `Chunk(items []int, size int) [][]int`:

- It splits `items` into consecutive chunks of length `size`. The last chunk may be shorter:
  `Chunk([1 2 3 4 5], 2)` → `[[1 2] [3 4] [5]]`.
- It returns `nil` if `size <= 0` or `items` is empty.
- **Appending to a returned chunk must never change `items`** or any other chunk.

Reference: https://gobyexample.com/slices
