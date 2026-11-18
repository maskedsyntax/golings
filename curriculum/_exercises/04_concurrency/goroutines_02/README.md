# goroutines_02: Results from goroutines

`MapParallel` applies `f` to every item, each in its own goroutine, and returns the results in
the **same order** as the input. Each goroutine writes only to `results[i]`, so no two
goroutines touch the same variable.

There is still a bug. Run the checks and look at what `MapParallel` returns.

## Your task

Fix `MapParallel` so it returns every result. The `race` check runs the tests under Go's race
detector (`go test -race`), which reports unsynchronized access to shared memory.

Reference: https://gobyexample.com/goroutines
