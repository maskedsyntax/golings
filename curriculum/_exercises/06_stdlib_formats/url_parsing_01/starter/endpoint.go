// url_parsing_01: Taking a URL apart
//
// Implement Parse. See README.md.
//
// Reference: https://gobyexample.com/url-parsing
package endpoint

import "errors"

// Endpoint is the part of a URL a client needs to connect.
type Endpoint struct {
	Scheme, Host, Port, Path, Token string
}

// Parse splits an http(s) URL into an Endpoint.
func Parse(raw string) (Endpoint, error) {
	// TODO
	return Endpoint{}, errors.New("not implemented")
}
