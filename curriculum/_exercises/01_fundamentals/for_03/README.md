# for_03: `break`, `continue` and labels

`continue` jumps to the next iteration. `break` exits a loop: by default the innermost one, or
a **labeled** outer loop with `break label`.

## Your task

Implement in `search.go`:

- `PositiveSum(xs []int) int` adds up only the positive numbers (use `continue`).
- `FirstPair(nums []int, target int) (i, j int, found bool)` finds the first pair of positions
  `i < j` with `nums[i] + nums[j] == target`. "First" means the smallest `i`, then the smallest
  `j`. If there's no such pair, it returns `0, 0, false`.

Reference: https://gobyexample.com/for
