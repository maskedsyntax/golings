// Mutant: hashes whatever arrived before the error.
package stream

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func HashReader(r io.Reader) (sum string, n int64, err error) {
	h := sha256.New()
	n, _ = io.Copy(h, r)
	return hex.EncodeToString(h.Sum(nil)), n, nil
}
