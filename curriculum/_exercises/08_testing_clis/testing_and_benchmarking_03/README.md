# testing_and_benchmarking_03: Benchmarks

Benchmarks measure how fast code runs. `go test -bench .` runs every `BenchmarkXxx` function:

```go
func BenchmarkSomething(b *testing.B) {
	for b.Loop() {
		Something()
	}
}
```

`join.go` (provided) has two ways to join strings: `JoinPlus` (`+=` in a loop) and
`JoinBuilder` (`strings.Builder`).

## Your task

Complete `BenchmarkJoin` in `join_test.go` so it runs two sub-benchmarks with `b.Run`:
`"plus"`, calling `JoinPlus(parts)` in a `b.Loop()` loop, and `"builder"`, calling
`JoinBuilder(parts)`.

golings runs your benchmark briefly and counts the calls. Afterwards, try
`go test -bench . -benchmem` in your workspace to see the difference.

Reference: https://gobyexample.com/testing-and-benchmarking
