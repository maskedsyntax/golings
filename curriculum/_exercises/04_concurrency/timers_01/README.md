# timers_01: A cancellable reminder

`time.Sleep` waits unconditionally. A **timer** fires once in the future and can be stopped
before it does:

```go
timer := time.NewTimer(2 * time.Second)
stopped := timer.Stop() // true if the timer hadn't fired yet
```

`time.AfterFunc(d, f)` is a timer that calls `f` (in its own goroutine) when it fires.

## Your task

Implement `Schedule(d time.Duration, f func()) (cancel func() bool)`:

- `f` runs once, after `d`, unless `cancel` is called first.
- `cancel` reports whether it prevented `f` from running (`false` if `f` has already started).

The starter sleeps in a goroutine, which can't be cancelled.

Reference: https://gobyexample.com/timers
