# channel_directions_03: Let the compiler find it

`DoubleAll` hands numbers to a `worker` goroutine and collects the doubled results. It hangs.
The bug is a single wrong channel name, the kind of mistake that channel directions turn
into a compile error.

## Your task

1. Give `worker`'s channel parameters directions (the `ast` check requires it).
2. Read the compiler error that appears.
3. Fix the bug.

Reference: https://gobyexample.com/channel-directions
