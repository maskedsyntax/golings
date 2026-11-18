# sha256_hashes_02: Hashing a stream

Files can be far bigger than memory. Because `sha256.New()` returns an `io.Writer`, you can
stream data into it a piece at a time.

## Your task

Implement `HashReader(r io.Reader) (sum string, n int64, err error)`. It returns the hex SHA-256
of everything read from `r` and the number of bytes read. A read error must be returned, not
hashed over.

Reference: https://gobyexample.com/sha256-hashes
