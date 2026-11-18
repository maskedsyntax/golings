# rate_limiting_03: A token bucket on paper

Production rate limiters (such as `golang.org/x/time/rate`) don't run a goroutine per limiter.
They compute the tokens from **elapsed time** on each request:

```text
tokens += (now − last) × rate
```

This bucket takes `now` as a parameter, so tests can drive it with a fake clock and nothing
depends on real timing. After a long quiet period it lets through far more requests than it
should.

## Your task

Find and fix the bug in `Allow`.

Reference: https://gobyexample.com/rate-limiting
