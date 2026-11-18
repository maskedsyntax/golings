# strings_and_runes_03: Truncate without breaking characters

`s[i:j]` slices a string by **byte** offsets. Cut in the middle of a multi-byte character
and you get invalid UTF-8, which prints as `�`.

## Your task

Implement `Truncate(s string, max int) string`, which shortens `s` to at most `max`
**characters**:

- If `s` has `max` characters or fewer, return it unchanged.
- Otherwise, return the first `max-1` characters followed by `…` (U+2026, a single character),
  so the result has exactly `max` characters.
- If `max <= 0`, return `""`.

```go
Truncate("hello world", 6) // "hello…"
Truncate("crème brûlée", 4) // "crè…"
```

Reference: https://gobyexample.com/strings-and-runes
