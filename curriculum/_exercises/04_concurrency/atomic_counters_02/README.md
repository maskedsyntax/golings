# atomic_counters_02: Compare and swap

Not every update is an `Add`. To keep the **maximum** of many concurrently observed values, you
have to read, compare and write, and another goroutine can sneak in between those steps.

`CompareAndSwap(old, new)` stores `new` **only if** the value is still `old`, and reports
whether it did:

```go
for {
	cur := v.Load()
	if /* no change needed */ || v.CompareAndSwap(cur, next) {
		break
	}
	// someone else changed it: try again with the new value
}
```

## Your task

Implement `Max.Observe` and `Max.Value`. `Observe` may be called from many goroutines at once.
The zero `Max` reports 0, and all observed values are non-negative.

Reference: https://gobyexample.com/atomic-counters
