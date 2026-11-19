# temporary_files_and_directories_02: A scratch directory

`os.MkdirTemp` creates a fresh directory with a unique name, which is handy for work that
produces many intermediate files.

## Your task

Implement `Scratch(fn func(dir string) error) error`:

- Create a temporary directory whose name starts with `scratch-`.
- Call `fn` with it and return `fn`'s error.
- Afterwards delete the directory **and everything `fn` put in it**, even if `fn` returns an
  error or panics.

Reference: https://gobyexample.com/temporary-files-and-directories
