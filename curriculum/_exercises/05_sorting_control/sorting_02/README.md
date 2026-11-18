# sorting_02: The median

The **median** of a list of numbers is the middle value once they are sorted. With an even
number of values, it's the average of the two in the middle.

```text
Median([3, 1, 2])     = 2
Median([4, 1, 3, 2])  = 2.5
```

## Your task

Implement `Median(nums []float64) (float64, bool)`. It returns `false` for an empty slice, and
it must not reorder the caller's slice.

Reference: https://gobyexample.com/sorting
