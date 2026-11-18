# channel_synchronization_02: Relay race

`Relay` runs every step **in its own goroutine**, and the steps must still run strictly one
after another, like runners passing a baton. Step `i` may begin only after step `i-1` has
finished. `Relay` returns when the last step is done.

Use channels as the baton. No mutexes, no WaitGroup, and no running the steps directly in
`Relay`'s goroutine.

The steps in the tests share a slice with no lock. That is only safe because each step
finishes before the next begins, and the race detector checks this.

Reference: https://gobyexample.com/channel-synchronization
