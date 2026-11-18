# base64_encoding_01: URL-safe tokens

`encoding/base64` comes in several flavours:

| encoding | alphabet | padding |
|---|---|---|
| `StdEncoding` | `A-Z a-z 0-9 + /` | `=` |
| `URLEncoding` | `A-Z a-z 0-9 - _` | `=` |
| `RawURLEncoding` | `A-Z a-z 0-9 - _` | none |

Tokens that go into URLs must not contain `+`, `/` or `=`.

## Your task

Fix `EncodeToken` and `DecodeToken` to use unpadded, URL-safe base64.

Reference: https://gobyexample.com/base64-encoding
