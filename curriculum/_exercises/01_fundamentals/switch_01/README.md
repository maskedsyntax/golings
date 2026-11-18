# switch_01: No fallthrough

In Go, each `case` ends by itself; there's no need for `break`. That also means an empty
`case` does **not** fall through into the next one.

## Your task

`IsWeekend("Saturday")` returns `false`. Fix `weekend.go` so Saturday and Sunday are
weekend days and every other input is not.

Reference: https://gobyexample.com/switch
