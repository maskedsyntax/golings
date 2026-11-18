# string_formatting_02: Human-readable sizes

`ls -h` and friends print sizes like `1.5 KiB`. Each unit is 1024 times the one before:
B, KiB, MiB, GiB, TiB.

## Your task

Implement `FormatBytes(n int64) string`:

| n | result |
|---|---|
| 0 | `0 B` |
| 1023 | `1023 B` |
| 1024 | `1.0 KiB` |
| 1536 | `1.5 KiB` |
| 10 << 20 | `10.0 MiB` |
| 1 << 30 | `1.0 GiB` |
| 1 << 50 | `1024.0 TiB` (TiB is the biggest unit) |

Bytes are whole numbers. Every other unit has exactly one decimal place.

Reference: https://gobyexample.com/string-formatting
