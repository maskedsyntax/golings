# epoch_02: Time to live

A token says when it expires as a Unix timestamp in **seconds**. `TTL` should say how long it
has left:

```go
TTL(1090, time.Unix(1000, 0))          // 1m30s
TTL(1010, time.Unix(1000, 500_000_000)) // 9.5s
TTL(900,  time.Unix(1000, 0))          // 0 (already expired)
```

It returns tiny durations like `90ns` instead.

## Your task

Fix `TTL`.

Reference: https://gobyexample.com/epoch
