# random_numbers_02: Random samples

`rand.Rand` has `Shuffle` and `Perm` for random orderings as well as single numbers.

## Your task

Implement `Sample(r *rand.Rand, items []string, k int) []string`. It returns `k` items chosen
at random, without repeats:

- Each item is equally likely to be picked.
- The caller's slice must not be changed.
- If `k >= len(items)`, return all the items (in random order). If `k <= 0`, return an empty
  slice.
- Use only `r` for randomness, so a fixed seed gives a fixed result.

Reference: https://gobyexample.com/random-numbers
