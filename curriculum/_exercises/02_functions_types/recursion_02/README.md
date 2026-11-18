# recursion_02: Towers of Hanoi

`n` disks of different sizes sit on peg `from`, largest at the bottom. Move them all to
peg `to`, one disk at a time, never putting a larger disk on a smaller one. A third peg,
`via`, is available as a spare.

The puzzle has an elegant recursive solution: if you can move `n-1` disks, you can move `n`.

## Your task

Implement `Hanoi(n int, from, to, via string) []string`. It returns the moves, in order, as
`"<from>-><to>"` strings, using the fewest possible moves (`2ⁿ - 1`).

```go
Hanoi(2, "A", "C", "B") // ["A->B", "A->C", "B->C"]
```

For `n == 0` there are no moves. The tests replay your moves and check that every one is
legal.

Reference: https://gobyexample.com/recursion
