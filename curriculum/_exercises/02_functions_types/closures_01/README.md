# closures_01: A counter that remembers

A **closure** is a function literal that refers to variables from the scope it was created
in. The variable lives on as long as the closure does, and each closure created by a new
call gets its own copy.

## Your task

Implement `Counter() func() int`. Each call to `Counter` returns a new function. Calling
that function returns 1, then 2, then 3, and so on. Two counters must count independently.

```go
a, b := Counter(), Counter()
a() // 1
a() // 2
b() // 1
```

Reference: https://gobyexample.com/closures
