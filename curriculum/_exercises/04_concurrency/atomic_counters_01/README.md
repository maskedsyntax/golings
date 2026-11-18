# atomic_counters_01: Counting without losing

`sync/atomic` performs simple operations on numbers **atomically**: no other goroutine can see
or interfere with a half-done update.

```go
var ops atomic.Uint64
ops.Add(1)
fmt.Println(ops.Load())
```

`Counter` is used from many goroutines at once, and it loses counts.

## Your task

Make `Counter` safe for concurrent use with `sync/atomic` (no mutex).

Reference: https://gobyexample.com/atomic-counters
