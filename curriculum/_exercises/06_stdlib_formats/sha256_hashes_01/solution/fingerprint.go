// sha256_hashes_01: Sum is not a hash function (reference solution)
package fingerprint

import (
	"crypto/sha256"
	"encoding/hex"
)

// Fingerprint returns the hex-encoded SHA-256 hash of data.
func Fingerprint(data []byte) string {
	sum := sha256.Sum256(data) // [32]byte
	return hex.EncodeToString(sum[:])
}
