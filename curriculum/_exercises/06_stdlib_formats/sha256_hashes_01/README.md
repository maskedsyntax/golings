# sha256_hashes_01: Sum is not a hash function

A SHA-256 hash is a 32-byte fingerprint of some data, usually written as 64 hex digits.

```go
sum := sha256.Sum256([]byte("abc"))  // [32]byte
hex.EncodeToString(sum[:])          // "ba7816bf...f20015ad"
```

The streaming API is `h := sha256.New()`, then `h.Write(...)` one or more times, then
`h.Sum(nil)`.

`Fingerprint` returns something far too long, and it's the same for every input apart from a
prefix.

## Your task

Fix `Fingerprint` so it returns the lowercase hex SHA-256 of `data`.

Reference: https://gobyexample.com/sha256-hashes
