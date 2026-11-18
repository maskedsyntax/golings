# url_parsing_02: Building query strings

`SearchURL` builds a search link by gluing strings together. That breaks as soon as the query
contains `&`, `+`, `%` or `#`, and it adds a second `?` when the base URL already has
parameters.

```go
SearchURL("https://example.com/search?lang=en", "go & rust", 2)
// should keep lang=en and set q="go & rust" and page="2"
```

`url.Values` collects parameters, and its `Encode` method escapes them properly.

## Your task

Fix `SearchURL`. Keep the base URL's existing parameters, set `q` and `page`, and return an
error if `base` doesn't parse.

Reference: https://gobyexample.com/url-parsing
