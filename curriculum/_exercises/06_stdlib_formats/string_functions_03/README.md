# string_functions_03: Whole words

`CountWord(text, word)` should count how many times `word` appears in `text` **as a whole
word**, ignoring case:

```go
CountWord("Go go GO! gopher, ago; go.", "go") // 4, not 6
```

`gopher` and `ago` contain "go" but are different words. An empty `word` matches nothing.

## Your task

Fix `CountWord`.

Reference: https://gobyexample.com/string-functions
