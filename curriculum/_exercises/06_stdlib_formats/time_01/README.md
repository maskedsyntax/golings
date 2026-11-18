# time_01: Comparing instants

A `time.Time` is an instant together with a location (time zone) used for display. These two
values are the same moment:

```go
utc := time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC)
cet := time.Date(2024, 3, 5, 14, 0, 0, 0, time.FixedZone("UTC+2", 2*60*60))
```

Compare instants with `Equal`, `Before` and `After`, not `==` or `<`.

## Your task

Fix both functions:

- `SameMoment` reports whether two times are the same instant, whatever their zones.
- `Overlaps` reports whether two bookings share any time. A booking that ends at 10:00 and
  one that starts at 10:00 do **not** overlap.

Reference: https://gobyexample.com/time
