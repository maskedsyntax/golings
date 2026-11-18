# tickers_01: Polling on a schedule

A **ticker** fires repeatedly at a fixed interval:

```go
ticker := time.NewTicker(500 * time.Millisecond)
defer ticker.Stop()
for t := range ticker.C {
	fmt.Println("tick at", t)
}
```

`Poll` calls `check` until it returns `true` and reports how many calls it took. The starter
calls `check` in a tight loop, hammering whatever it's checking.

## Your task

Make `Poll` wait one `interval` (on a ticker) **before each call** to `check`, and stop the
ticker when it returns.

Reference: https://gobyexample.com/tickers
