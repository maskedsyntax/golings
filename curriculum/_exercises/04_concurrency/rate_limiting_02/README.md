# rate_limiting_02: Allowing bursts

Sometimes short **bursts** are fine as long as the average rate stays limited. A bursty limiter
holds up to `burst` tokens, gains one token per `every`, and each request spends one.

## Your task

Implement the `Limiter`:

- `NewLimiter(burst, every)` starts **full**, so the first `burst` calls to `Wait` return at once.
- `Wait()` blocks until a token is available and takes it.
- A token is added every `every`, but the limiter never holds more than `burst`, so a long
  idle period doesn't allow an unlimited burst.
- `Stop()` stops the refilling goroutine.

A buffered channel makes a natural token bucket.

Reference: https://gobyexample.com/rate-limiting
