// Mutant: handles padding but not the URL alphabet.
package decode

import (
	"encoding/base64"
	"strings"
)

func DecodeAny(s string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(strings.TrimRight(s, "="))
}
