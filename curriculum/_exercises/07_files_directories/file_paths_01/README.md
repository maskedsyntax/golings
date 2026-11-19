# file_paths_01: Portable paths

`path/filepath` builds and splits file paths with the separator of the operating system you're
running on: `/` on Linux and macOS, `\` on Windows.

```go
filepath.Join("dir1", "dir2", "file")  // dir1/dir2/file  (or dir1\dir2\file)
filepath.Dir("dir1/file")              // dir1
filepath.Base("dir1/file")             // file
```

`BackupPath` builds paths by hand. It panics on a bare file name, and it would produce mixed
separators on Windows.

## Your task

Fix `BackupPath`: `docs/notes.txt` becomes `docs/backup/notes.txt.bak`, and `notes.txt`
becomes `backup/notes.txt.bak`.

Reference: https://gobyexample.com/file-paths
