// base64_encoding_01: URL-safe tokens
//
// The tokens contain + / and =. Fix them. See README.md.
//
// Reference: https://gobyexample.com/base64-encoding
package token

import "encoding/base64"

// EncodeToken encodes b as unpadded URL-safe base64.
func EncodeToken(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// DecodeToken reverses EncodeToken.
func DecodeToken(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
