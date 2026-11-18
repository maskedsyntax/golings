# interfaces_03: A nil that isn't nil

An interface value has two parts: a dynamic **type** and a **value**. It equals `nil` only
when both are missing. A nil pointer stored in an interface still has a type, so the
interface is **not** nil:

```go
var p *Circle      // nil pointer
var s Shape = p    // s != nil, because s holds the type *Circle
```

## Your task

`Find(name)` returns the shape with that name, or `nil` if there isn't one. Callers check
`if s := Find(name); s != nil { … }`, but for unknown names that check passes, and the
program crashes. Fix `Find`.

Reference: https://gobyexample.com/interfaces
