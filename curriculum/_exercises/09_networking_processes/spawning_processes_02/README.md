# spawning_processes_02: Pipes and stderr

`Filter` runs a command, feeds it `input` on standard input, and returns its standard output.
For a program like `tr a-z A-Z` that reads until the end of its input, it hangs until the
context kills it. When the command fails, the error says only `exit status 3`, which doesn't
help anyone.

## Your task

Fix `Filter`:

- The child must see the end of its input.
- If the command fails, the error must include what the child wrote to standard error.

Reference: https://gobyexample.com/spawning-processes
