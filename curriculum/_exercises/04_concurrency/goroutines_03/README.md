# goroutines_03: A shared total

`ParallelSum` adds up several chunks of numbers, one goroutine per chunk. It waits for every
goroutine, and it usually even returns the right answer. Run the checks: the race detector
disagrees.

When two goroutines read and write the same variable without synchronization, that is a
**data race**. Updates can be lost, and in Go a racy program has no guaranteed behaviour at all.

## Your task

Make `ParallelSum` race-free without a mutex. Each goroutine should write only to its own
variable, and you combine them after waiting. (Mutexes come later in this module.)

Reference: https://gobyexample.com/goroutines
