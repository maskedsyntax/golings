// Mutant: only half of the hash.
package fingerprint

import (
	"crypto/sha256"
	"encoding/hex"
)

func Fingerprint(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:16])
}
