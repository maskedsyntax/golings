# multiple_return_values_02: Minimum, maximum, and ok

A common Go idiom is to return an extra boolean, usually called `ok`, that says whether
the other results are valid. You'll see it again with map lookups and type assertions.

## Your task

Implement `MinMax(nums []int) (lo, hi int, ok bool)`:

- It returns the smallest and largest element of `nums`, and `ok == true`.
- For an empty (or nil) slice it returns `0, 0, false`.

Reference: https://gobyexample.com/multiple-return-values
