# reading_files_02: Reading part of a file

For big files you often want only a piece. `os.Open` gives you an `*os.File`, which can read
at any position:

```go
f, err := os.Open(path)
...
defer f.Close()
n, err := f.ReadAt(buf, 100) // len(buf) bytes starting at byte 100
```

The exercise comes with `data/alphabet.txt` (`a` to `z` and a newline).

## Your task

Implement `Peek(path string, offset int64, n int) (string, error)`. It returns up to `n` bytes
starting at `offset`:

- Near the end of the file, return the bytes that exist.
- An offset past the end gives `""` with no error.
- A missing file or a negative offset is an error.

Reference: https://gobyexample.com/reading-files
