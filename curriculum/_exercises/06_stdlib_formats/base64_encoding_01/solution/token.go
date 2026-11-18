// base64_encoding_01: URL-safe tokens (reference solution)
package token

import "encoding/base64"

// EncodeToken encodes b as unpadded URL-safe base64.
func EncodeToken(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

// DecodeToken reverses EncodeToken.
func DecodeToken(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
