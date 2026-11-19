# writing_files_01: Flush before you close

`bufio.Writer` batches small writes into larger ones, which is much faster. It holds the data
in memory until its buffer is full, or until you call `Flush`.

`WriteLines` writes nothing at all for short inputs, and loses the end of long ones.

## Your task

Fix `WriteLines` so every line reaches the file, each followed by `\n`. Report write and close
errors.

Reference: https://gobyexample.com/writing-files
