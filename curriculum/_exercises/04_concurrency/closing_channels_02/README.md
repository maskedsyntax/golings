# closing_channels_02: Many senders, one close

`SquareAll` shares the work among several worker goroutines that all send on one `results`
channel. The caller ranges over `results` until it is closed. Run the checks: the program
panics.

Closing a channel twice panics, and so does sending on a closed channel. With many senders,
none of them knows whether it is the last one.

## Your task

Fix `SquareAll` so `results` is closed exactly once, after every worker has finished.

Reference: https://gobyexample.com/closing-channels
