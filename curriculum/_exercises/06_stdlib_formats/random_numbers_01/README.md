# random_numbers_01: Rolling dice

`math/rand/v2` generates pseudo-random numbers. Tests pass in a generator with a fixed *seed*,
so the "random" results are the same on every run:

```go
r := rand.New(rand.NewPCG(1, 2)) // seeded: reproducible
r.IntN(10)                       // 0 <= n < 10
```

`Roll` should return `n` rolls of a six-sided die, but its rolls include 0 and never 6.

## Your task

Fix `Roll`.

Reference: https://gobyexample.com/random-numbers
