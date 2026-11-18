# mutexes_01: Guarding a map

For state more complex than a single number, use a **mutex** so that only one goroutine at a
time touches it:

```go
c.mu.Lock()
defer c.mu.Unlock()
c.counters[name]++
```

`Container` already has a mutex, but nothing uses it.

## Your task

Make `Inc` and `Get` safe to call from many goroutines at once.

Reference: https://gobyexample.com/mutexes
