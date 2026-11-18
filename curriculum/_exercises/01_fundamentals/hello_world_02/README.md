# hello_world_02: Imports and exported names

Two rules Go enforces at compile time:

1. Every imported package must be used, and every package you use must be imported.
2. A name is **exported** (usable from other packages) only if it starts with a capital letter.

## Your task

Fix `main.go` so it compiles and prints exactly:

```text
Go has exported names.
They start with a capital letter.
```

Reference: https://gobyexample.com/hello-world
