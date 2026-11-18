# closures_03: Memoize a function

Closures can keep private state between calls. Here that state is a cache.

## Your task

Implement `Memoize(f func(int) int) func(int) int`. The returned function gives the same
answers as `f`, but calls `f` at most **once per distinct argument** and remembers the result.

- Each call to `Memoize` gets its own cache. Memoizing two different functions must never mix
  up their answers.
- Results of `0` are valid and must be cached too.

Reference: https://gobyexample.com/closures
