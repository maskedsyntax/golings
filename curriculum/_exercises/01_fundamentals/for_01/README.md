# for_01: Go has only `for`

There's no `while` in Go. `for` does every kind of loop:

```go
for i := 0; i < 3; i++ { }   // classic
for n > 0 { }                // "while"
for { }                      // forever
```

## Your task

Fix `loops.go` so it compiles:

- `CountDown(n int) []int` returns `n, n-1, …, 1`. It returns an empty result for `n <= 0`.
- `SumTo(n int) int` returns `1 + 2 + … + n`, or 0 for `n <= 0`.

Reference: https://gobyexample.com/for
