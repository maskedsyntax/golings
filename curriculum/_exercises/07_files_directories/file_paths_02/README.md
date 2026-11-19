# file_paths_02: Extensions and relative paths

## Your task

Implement two helpers:

- `ChangeExt(p, ext string) string` replaces the extension of the **file name** in `p`:
  `a/b.txt` → `a/b.md`, and `archive.tar.gz` → `archive.tar.zip`. A dot in a directory name
  doesn't count: `a.b/c` → `a.b/c.txt`.
- `RelTo(base, target string) (string, error)` returns `target` relative to `base`. For
  `/srv/app` and `/srv/data/x` that's `../data/x`. It's an error when there's no relative path,
  for example from a relative `base` to an absolute `target`.

Reference: https://gobyexample.com/file-paths
