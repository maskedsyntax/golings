# time_02: Business days

`time` does calendar arithmetic for you: `t.AddDate(years, months, days)` handles month
lengths and leap years, and `t.Weekday()` says which day of the week it is.

## Your task

Implement `AddBusinessDays(t time.Time, n int) time.Time`. It returns the time `n` working
days after `t`, skipping Saturdays and Sundays and keeping the time of day.

- Friday + 1 is Monday.
- Thursday + 5 is the next Thursday.
- Saturday + 1 is Monday.
- `n == 0` returns `t` unchanged, even on a weekend.

`n` is never negative.

Reference: https://gobyexample.com/time
