# timeouts_02: An overall deadline

`CollectFor` should gather values from `ch` for **at most** `d` in total, then return what it
has. When values arrive steadily, it never returns.

## Your task

Fix `CollectFor` so `d` limits the whole call, not the wait for each single value.

Reference: https://gobyexample.com/timeouts
