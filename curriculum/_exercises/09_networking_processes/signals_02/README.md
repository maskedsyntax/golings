# signals_02: Shutting down on SIGTERM

`signal.Notify(ch, sigs...)` asks Go to deliver those signals to `ch` instead of killing the
program. Container runtimes and `kill` send `SIGTERM`, and Ctrl+C sends `SIGINT`
(`os.Interrupt`). A well-behaved program cleans up on either.

`Run(ready, cleanup)` sets up signal handling, calls `ready`, waits for a signal, and calls
`cleanup`. The tests start your code in a real child process and send it a real `SIGTERM`.

## Your task

Fix `Run` so both `SIGINT` and `SIGTERM` lead to `cleanup` and a normal return. `go vet`
reports the first problem.

Signals work this way only on Unix-like systems, so this exercise is checked on Linux and macOS.

Reference: https://gobyexample.com/signals
