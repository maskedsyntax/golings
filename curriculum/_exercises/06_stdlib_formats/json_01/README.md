# json_01: Exported fields and tags

`json.Marshal` turns a Go value into JSON. For a struct, it writes each **exported** field, and
a *struct tag* chooses the key name:

```go
type Point struct {
	X int `json:"x"`
	Y int `json:"y,omitempty"` // left out when zero
}
```

`Encode` should produce

```json
{"name":"Ada","email":"ada@example.com","admin":true}
```

leaving `admin` out when it's false. Right now it produces `{}`.

## Your task

Fix the `User` type. `NewUser` and `Encode` don't need to change.

Reference: https://gobyexample.com/json
