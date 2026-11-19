# line_filters_02: A tiny grep

## Your task

Implement `Grep(r io.Reader, w io.Writer, pattern string) (int, error)`. Like `grep -n`, it:

- writes each line that contains `pattern` as `<line number>:<line>`, with line numbers
  starting at 1,
- returns how many lines matched,
- handles Windows line endings (`\r\n`): the `\r` must not appear in the output.

```text
input:   alpha\r\nbeta\r\ngamma alpha\n
pattern: alpha
output:  1:alpha\n3:gamma alpha\n   (2 matches)
```

Reference: https://gobyexample.com/line-filters
