# struct_embedding_02: Overriding and delegating

If the outer struct declares a method with the same name as a promoted one, the outer method
wins. It can still call the embedded one explicitly, through the embedded field, whose
name is the type name.

## Your task

`Service` embeds a `*Logger`. Several services may share one logger. Give `*Service` its own
`Log(msg string)` method that prefixes each line with the service name, then delegates to the
logger:

```go
l := &Logger{}
api := &Service{Logger: l, Name: "api"}
api.Log("started") // l.Lines() == ["[api] started"]
```

Everything else `Logger` provides, such as `Lines`, should still be promoted.

Reference: https://gobyexample.com/struct-embedding
