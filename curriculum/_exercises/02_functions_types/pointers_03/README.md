# pointers_03: Pointers to copies

`for _, v := range nums` copies each element into `v`. Taking `&v` gives you a pointer to
that copy, not to the element in the slice.

## Your task

`ElementPointers(nums []int) []*int` should return a pointer to **each element of `nums`**,
so that writing through the i-th pointer changes `nums[i]`:

```go
nums := []int{1, 2, 3}
ptrs := ElementPointers(nums)
*ptrs[1] = 20 // nums is now [1 20 3]
```

Right now, writing through the pointers has no effect on `nums`. Fix it.

Reference: https://gobyexample.com/pointers
