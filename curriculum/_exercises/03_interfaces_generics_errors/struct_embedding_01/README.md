# struct_embedding_01: Promoted fields and methods

A struct can **embed** another type by listing it without a field name:

```go
type container struct {
	base         // embedded
	str string
}
```

The embedded type's fields and methods are **promoted**: `c.num` and `c.describe()` work
directly on a `container`. A promoted method also counts toward satisfying interfaces.

## Your task

The tests treat a `Container` as having an `ID` field and a `Describe` method, and pass it
where a `Describer` is expected. Change `Container` so that `Base` is embedded.

Reference: https://gobyexample.com/struct-embedding
