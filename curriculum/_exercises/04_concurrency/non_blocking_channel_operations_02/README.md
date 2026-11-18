# non_blocking_channel_operations_02: Slow subscribers

`Hub` delivers each published value to every subscriber channel. Some subscribers are slow,
or have stopped reading altogether. One stuck subscriber freezes `Publish`, and with it every
other subscriber.

## Your task

Make `Publish` never block. When a subscriber's buffer is full, drop the value for **that
subscriber only** and add one to `h.Dropped`.

Reference: https://gobyexample.com/non-blocking-channel-operations
