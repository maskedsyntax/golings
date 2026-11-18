# arrays_01: Arrays are values

An array type includes its length (`[3]int`), and arrays are **values**: assigning one copies
every element. Slices share their backing array instead.

## Your task

`History(moves []int) [][3]int` simulates a three-cell board. Move `i` writes `moves[i]` into
cell `i % 3`, and after each move a snapshot of the board is recorded:

```text
History([]int{7, 8, 9, 1}) == [[7 0 0] [7 8 0] [7 8 9] [1 8 9]]
```

The current code records the same snapshot over and over. Fix it.

Reference: https://gobyexample.com/arrays
