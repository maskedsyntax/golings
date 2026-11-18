# range_over_channels_01: Ranging until closed

`for … range` works on channels too. It receives values until the channel is closed:

```go
for elem := range queue {
	fmt.Println(elem)
}
```

`Sum` assumes the sender sends exactly three numbers. That is only right some of the time.

## Your task

Rewrite `Sum` with `for … range` (the `ast` check requires it) so it adds up every value until
the channel is closed.

Reference: https://gobyexample.com/range-over-channels
