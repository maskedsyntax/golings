// sha256_hashes_02: Hashing a stream
//
// Implement HashReader. See README.md.
//
// Reference: https://gobyexample.com/sha256-hashes
package stream

import (
	"errors"
	"io"
)

// HashReader returns the hex SHA-256 of r's contents and how many bytes it read.
func HashReader(r io.Reader) (sum string, n int64, err error) {
	// TODO
	return "", 0, errors.New("not implemented")
}
