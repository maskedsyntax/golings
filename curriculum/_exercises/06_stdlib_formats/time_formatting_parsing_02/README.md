# time_formatting_parsing_02: Two log formats

Logs from different tools write times differently:

```text
2024-03-05T10:00:00+02:00        RFC 3339 (time.RFC3339)
05/Mar/2024:10:00:00 +0200       web server access logs
```

Both of those are the same instant: 08:00 UTC.

## Your task

Implement `ParseLogTime(s string) (time.Time, error)`. It accepts either format (RFC 3339 may
have fractional seconds) and returns the time **in UTC**. Anything else is an error.

Reference: https://gobyexample.com/time-formatting-parsing
