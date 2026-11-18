// sha256_hashes_01: Sum is not a hash function
//
// Fingerprint misuses Sum. Fix it. See README.md.
//
// Reference: https://gobyexample.com/sha256-hashes
package fingerprint

import (
	"crypto/sha256"
	"encoding/hex"
)

// Fingerprint returns the hex-encoded SHA-256 hash of data.
func Fingerprint(data []byte) string {
	h := sha256.New()
	return hex.EncodeToString(h.Sum(data))
}
