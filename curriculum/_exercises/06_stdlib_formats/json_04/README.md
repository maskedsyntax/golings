# json_04: Decoding by kind

Some JSON says what type a value is in one field and keeps the details in another:

```json
[
  {"kind": "circle", "spec": {"r": 1}},
  {"kind": "rect",   "spec": {"w": 2, "h": 3}}
]
```

You can't decode `spec` until you've read `kind`. `json.RawMessage` holds a piece of JSON
undecoded, so you can decode it later once you know which type it is.

## Your task

Implement `DecodeShapes(data []byte) ([]Shape, error)`. It returns a `Circle` or a `Rect` for
each element, in order. An unknown `kind`, or a spec that doesn't fit its type, is an error.

Reference: https://gobyexample.com/json
