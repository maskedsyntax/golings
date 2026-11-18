# recursion_01: Base case and recursive case

A **recursive** function calls itself. To finish, it needs a **base case**: an input it
answers directly, without calling itself again.

## Your task

Implement, **recursively**:

- `Factorial(n int) int`: `n! = n × (n-1) × … × 1`, with `Factorial(0) == 1`.
- `Fib(n int) int`: the Fibonacci numbers `0, 1, 1, 2, 3, 5, 8, …`, so `Fib(0) == 0`,
  `Fib(1) == 1` and `Fib(n) == Fib(n-1) + Fib(n-2)`.

You can assume `n >= 0`. This exercise checks that both functions call themselves.

Reference: https://gobyexample.com/recursion
