# exit_01: os.Exit skips defers

`os.Exit(code)` ends the program immediately with an exit status. **Deferred functions don't
run**, and that includes the `Flush` that writes out buffered output.

This program should print

```text
checking 3 files...
found 2 problems
```

and exit with status 2 (the number of problems). It exits with the right status but prints
nothing.

## Your task

Fix it. A common idiom is `func main() { os.Exit(run()) }`, where `run` does the work and its
defers run before it returns.

Reference: https://gobyexample.com/exit
