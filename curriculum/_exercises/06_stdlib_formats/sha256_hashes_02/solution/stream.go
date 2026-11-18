// sha256_hashes_02: Hashing a stream (reference solution)
package stream

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// HashReader returns the hex SHA-256 of r's contents and how many bytes it read.
func HashReader(r io.Reader) (sum string, n int64, err error) {
	h := sha256.New()
	n, err = io.Copy(h, r) // streams; never holds all of r in memory
	if err != nil {
		return "", n, err
	}
	return hex.EncodeToString(h.Sum(nil)), n, nil
}
