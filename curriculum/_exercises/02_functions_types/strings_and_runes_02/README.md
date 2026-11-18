# strings_and_runes_02: Unicode-aware palindromes

The `unicode` package answers questions about runes in any script: `unicode.IsLetter`,
`unicode.IsDigit`, `unicode.ToLower`, and more.

## Your task

Implement `IsPalindrome(s string) bool`. A string is a palindrome if its letters and digits,
ignoring case, read the same forwards and backwards. Spaces, punctuation and other symbols
are ignored.

```go
IsPalindrome("A man, a plan, a canal: Panama") // true
IsPalindrome("Été")                            // true
IsPalindrome("étè")                            // false: é and è are different letters
```

Every letter counts, not just ASCII ones. The empty string is a palindrome.

Reference: https://gobyexample.com/strings-and-runes
