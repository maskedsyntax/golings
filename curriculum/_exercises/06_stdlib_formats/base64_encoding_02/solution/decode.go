// base64_encoding_02: Decoding whatever arrives (reference solution)
package decode

import (
	"encoding/base64"
	"strings"
)

var toStd = strings.NewReplacer("-", "+", "_", "/")

// DecodeAny decodes standard or URL-safe base64, with or without padding.
func DecodeAny(s string) ([]byte, error) {
	s = strings.TrimRight(s, "=")
	s = toStd.Replace(s)
	return base64.RawStdEncoding.DecodeString(s)
}
