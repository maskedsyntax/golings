# defer_03: Defer in a loop

`TotalSize` opens each file, reads it, and closes it with `defer`. It gives the right answer,
but deferred calls run when the **function** returns, not at the end of each loop iteration.
Every file stays open until `TotalSize` finishes. With enough files, the program runs out of
file handles.

## Your task

Restructure `TotalSize` so **at most one** file is open at a time, and every file (including
one whose read fails) is closed.

Reference: https://gobyexample.com/defer
