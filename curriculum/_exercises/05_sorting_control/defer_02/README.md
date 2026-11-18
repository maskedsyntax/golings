# defer_02: When defer looks

`Run` executes some steps and, when it returns, adds a cleanup line that names the **last**
step that ran:

```text
Run([a b c]) → [run a, run b, run c, cleanup after c]
```

It writes `cleanup after ` with no step name, and the cleanup line doesn't even end up in
the result.

Two facts about `defer` explain both bugs:

1. The arguments of a deferred call are evaluated **immediately**, when the `defer` statement
   runs.
2. A deferred function runs **after** the `return` statement has set the result, so it can
   only change results that have a name.

## Your task

Fix `Run`. Keep the cleanup in a `defer` (the `ast` check requires it).

Reference: https://gobyexample.com/defer
