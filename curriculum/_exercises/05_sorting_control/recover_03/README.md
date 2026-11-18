# recover_03: A panic in a goroutine

`RunJobs` runs each job in its own goroutine and is meant to survive jobs that panic:
`errs[i]` describes job `i`'s panic, and the other jobs still finish. It has a `recover`, but a
single panicking job still crashes the whole program.

## Your task

Fix `RunJobs` so every panic is recovered and reported in `errs[i]`, and jobs that don't panic
report `nil`.

Reference: https://gobyexample.com/recover
