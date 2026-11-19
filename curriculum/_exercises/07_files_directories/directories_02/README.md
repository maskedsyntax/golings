# directories_02: Walking a tree

`filepath.WalkDir` visits every file and directory under a root, calling your function for
each. Returning `filepath.SkipDir` from the function skips a directory's contents.

## Your task

Implement `TotalSize(root string) (int64, error)`. It returns the total size in bytes of every
regular file under `root`, at any depth, **except** inside hidden directories (names starting
with `.`, like `.git`).

The root itself is never skipped, even if it's `"."`.

Reference: https://gobyexample.com/directories
