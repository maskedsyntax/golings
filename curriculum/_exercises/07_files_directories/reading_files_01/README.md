# reading_files_01: Reading a whole file

`os.ReadFile` reads a whole file into memory, which is the simplest choice for small files:

```go
data, err := os.ReadFile("data/poem.txt")
```

This exercise comes with `data/poem.txt`. The checks read their own copy of the file, so
editing yours changes nothing.

## Your task

`CountLines` doesn't compile. Fix it, and return the error when the file can't be read.

Reference: https://gobyexample.com/reading-files
