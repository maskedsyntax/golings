# string_functions_02: Key=value pairs

Settings sometimes arrive as one line of `key=value` pairs separated by semicolons:

```text
name=Ada; lang = Go ;expr=a=b;; novalue ; empty=
```

## Your task

Implement `ParseKV(line string) map[string]string`:

- Pairs are separated by `;`.
- Each pair splits at its **first** `=`, so `expr=a=b` means key `expr` with value `a=b`.
- Keys and values have their surrounding spaces trimmed.
- Pieces with no `=`, and pieces with an empty key, are skipped.
- Always return a non-nil map.

The line above parses to `{"name": "Ada", "lang": "Go", "expr": "a=b", "empty": ""}`.

`strings.Cut` is made for this.

Reference: https://gobyexample.com/string-functions
