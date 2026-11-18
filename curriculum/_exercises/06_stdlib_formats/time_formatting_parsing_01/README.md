# time_formatting_parsing_01: Reference-time layouts

Go formats and parses times with a *layout*: the reference time

```text
Mon Jan 2 15:04:05 MST 2006
```

written the way you want your times to look. `2006` means the year, `01` the month, `02` the
day, `15` the 24-hour hour, `3` the 12-hour hour and `PM` the AM/PM marker.

## Your task

Fix the three layouts:

| function | example |
|---|---|
| `FormatDate` | `2024-03-05` |
| `ParseDate` | parses `"2024-03-05"` (UTC midnight) |
| `Clock` | `3:04 PM`, `9:05 AM` |

Reference: https://gobyexample.com/time-formatting-parsing
