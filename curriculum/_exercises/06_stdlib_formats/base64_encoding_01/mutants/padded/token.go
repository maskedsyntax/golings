// Mutant: URL alphabet but still padded.
package token

import "encoding/base64"

func EncodeToken(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

func DecodeToken(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
