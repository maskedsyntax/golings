# directories_01: Listing a directory

`os.ReadDir` lists a directory's entries (files and subdirectories), sorted by name. Each
`os.DirEntry` tells you its `Name()` and whether it `IsDir()`.

## Your task

Implement `ListGoFiles(dir string) ([]string, error)`. It returns the names of the Go source
files directly inside `dir`, sorted:

- only regular `.go` files, not directories,
- no `_test.go` files,
- no recursion into subdirectories.

A missing directory is an error.

Reference: https://gobyexample.com/directories
