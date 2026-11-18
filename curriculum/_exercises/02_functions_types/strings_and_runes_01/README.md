# strings_and_runes_01: Bytes are not characters

A Go string is a read-only sequence of **bytes**, usually UTF-8 text. A character such as
`é` (2 bytes) or `日` (3 bytes) takes several bytes. Go calls a decoded character a
**rune**.

## Your task

Both functions in `text.go` work for ASCII but break on other text:

- `Length(s string) int` should return the number of characters: `Length("héllo") == 5`.
- `Reverse(s string) string` should reverse the characters: `Reverse("añb") == "bña"`.

Fix them.

Reference: https://gobyexample.com/strings-and-runes
