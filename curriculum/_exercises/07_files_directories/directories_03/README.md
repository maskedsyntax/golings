# directories_03: Creating directories safely

`EnsureDirs(root, rels)` should create each relative directory path under `root`:

```go
EnsureDirs(root, []string{"logs", "data/cache/images"})
```

It has two problems:

1. It fails on nested paths whose parents don't exist yet, and when run a second time.
2. It trusts its input: `"../escape"` creates a directory **outside** `root`.

## Your task

Fix `EnsureDirs`:

- Create any missing parents.
- Succeed if the directories already exist.
- Return an error, without creating anything, for a path that isn't local to `root` (a `..`
  escape, an absolute path, or `""`).

Reference: https://gobyexample.com/directories
