# base64_encoding_02: Decoding whatever arrives

Different systems send base64 in different flavours: standard or URL alphabet, padded or not.

## Your task

Implement `DecodeAny(s string) ([]byte, error)`. It accepts any of the four combinations and
returns an error for input that isn't base64 at all.

```go
DecodeAny("+/+/")  // standard
DecodeAny("-_-_")  // URL-safe
DecodeAny("aGk=")  // padded  -> "hi"
DecodeAny("aGk")   // raw     -> "hi"
```

Reference: https://gobyexample.com/base64-encoding
