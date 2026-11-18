# worker_pools_02: Waiting in the wrong place

`CountLetters` uses a pool of workers to measure every word, and the workers send their
answers on a `results` channel. It hangs.

## Your task

Fix `CountLetters` so it collects every result while the pool is running, and keep the number
of workers it uses.

Reference: https://gobyexample.com/worker-pools
