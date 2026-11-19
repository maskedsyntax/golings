# writing_files_02: Atomic writes

If a program crashes, or its input fails, while it is overwriting a file, the file ends up
half old and half new. The standard fix is to write a temporary file in the same directory
and then **rename** it over the original. Readers then see either the whole old file or the
whole new one.

## Your task

Implement `WriteFileAtomic(path string, r io.Reader) error`. It replaces the file at `path`
with everything read from `r`:

- If reading `r` fails, return the error, leave the original file exactly as it was, and leave
  no temporary files behind.
- If `path` doesn't exist yet, create it.

Reference: https://gobyexample.com/writing-files
