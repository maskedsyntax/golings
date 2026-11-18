# regular_expressions_03: Replacement templates

`ReplaceAllString` can reuse what the groups captured: `$1` (or `${1}`) in the replacement
means "the text of group 1".

`Username` turns `"Lovelace, Ada"` into `"ada_lovelace"`. It returns only `"lovelace"`.

## Your task

Fix `Username`.

Reference: https://gobyexample.com/regular-expressions
