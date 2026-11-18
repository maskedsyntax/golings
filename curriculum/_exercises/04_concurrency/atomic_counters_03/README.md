# atomic_counters_03: Close exactly once

`Conn.Close` should release the connection the **first** time it is called and return
`ErrClosed` on every later call, even when many goroutines call it at the same moment. Releasing
twice would corrupt a real connection pool.

## Your task

Make `Close` safe for concurrent use with `sync/atomic` (no mutex): `release` must run exactly
once, and exactly one caller gets `nil`.

Reference: https://gobyexample.com/atomic-counters
