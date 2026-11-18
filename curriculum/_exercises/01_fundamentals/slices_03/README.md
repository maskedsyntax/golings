# slices_03: Remove duplicate values

## Your task

Implement `Unique(numbers []int) []int`. It returns the distinct numbers **in the order they
first appear**:

```text
Unique([1 2 1 3 2]) == [1 2 3]
```

Empty or nil input returns an empty, **non-nil** slice. Don't modify the input.

Tip: the standard library's `slices` package has helpers such as `slices.Contains`.

Reference: https://gobyexample.com/slices
