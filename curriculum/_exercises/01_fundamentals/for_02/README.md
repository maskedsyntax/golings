# for_02: Collatz steps

Start with a positive integer `n`. If it's even, halve it; if it's odd, replace it with
`3n + 1`. Repeat until you reach 1.

## Your task

Implement `Steps(n int) int`, which returns how many transformations it takes to reach 1:

- `Steps(1)` is `0`, because no transformation is needed.
- `Steps(6)` is `8`: 6 → 3 → 10 → 5 → 16 → 8 → 4 → 2 → 1.
- For `n <= 0`, return `-1`.

Reference: https://gobyexample.com/for
