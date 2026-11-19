# temporary_files_and_directories_01: Cleaning up a temp file

`os.CreateTemp(dir, pattern)` creates a new file with a unique name: `"report-*.txt"` becomes
something like `report-123456.txt`, in the system temp directory when `dir` is `""`. The file
stays there until **you** remove it.

`WithTempFile` writes data to a temp file and passes its path to `fn`. It never closes or
deletes the file, so every call leaks one.

## Your task

Fix `WithTempFile` so the file is closed before `fn` runs and removed afterwards, whether `fn`
succeeds or fails. Keep returning `fn`'s error.

Reference: https://gobyexample.com/temporary-files-and-directories
