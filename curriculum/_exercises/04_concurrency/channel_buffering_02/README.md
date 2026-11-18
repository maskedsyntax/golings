# channel_buffering_02: The first answer wins

`First` asks several sources the same question at once and returns whichever answer arrives
first. It returns the right answer, but it leaks goroutines. Every source that loses the race
stays blocked forever, trying to send an answer nobody will receive.

In a long-running server a leak like this grows until the process runs out of memory.

## Your task

Fix the leak. After `First` returns and the slower sources finish, no goroutines from `First`
may remain. The test counts goroutines with `runtime.NumGoroutine`.

Reference: https://gobyexample.com/channel-buffering
