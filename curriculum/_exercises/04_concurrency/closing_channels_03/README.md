# closing_channels_03: Stopping every worker

`Start` launches `n` workers that keep calling `work` until they are told to stop. It returns
a `stop` function. Calling `stop` must:

1. tell **every** worker to stop, and
2. return only once all of them have exited, so `work` is never called after `stop` returns.

Right now `stop` hangs.

## Your task

Fix `Start`. A closed channel works as a broadcast signal.

Reference: https://gobyexample.com/closing-channels
