# waitgroups_03: Collecting every error

`RunAll` runs a set of tasks and reports **every** failure, not just the first. Right now it
runs them one at a time.

## Your task

Make `RunAll` run all tasks concurrently and return `errors.Join` of their errors, **in task
order**. It returns `nil` when every task succeeds.

Reference: https://gobyexample.com/waitgroups
