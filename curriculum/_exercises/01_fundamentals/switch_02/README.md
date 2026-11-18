# switch_02: A switch with no expression

`switch { case cond1: …; case cond2: … }` checks conditions top to bottom and runs the first
match. It's often clearer than a long `if`/`else if` chain.

## Your task

Implement `FizzBuzz(n int) string` **using a switch**:

- `"FizzBuzz"` if `n` is divisible by both 3 and 5,
- `"Fizz"` if it's divisible by 3,
- `"Buzz"` if it's divisible by 5,
- otherwise the number itself, as text (`"7"`).

This exercise checks that `FizzBuzz` uses a `switch` statement.

Reference: https://gobyexample.com/switch
