# reading_files_03: Very long lines

`bufio.Scanner` is the usual way to read a file line by line. It has a limit, though: by default
a line (a *token*) can be at most 64 KiB. When a line is longer, `Scan` stops and returns
false, exactly as it does at the end of the file. The only way to tell the two apart is
`sc.Err()`.

`LongestLine` reports 5 for a file whose longest line has 200,000 characters.

## Your task

Fix `LongestLine`:

- Lines up to 1 MiB (1 << 20 bytes) must work.
- A longer line is an error, and so is any other read error.

Reference: https://gobyexample.com/reading-files
