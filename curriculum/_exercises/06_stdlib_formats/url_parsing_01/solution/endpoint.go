// url_parsing_01: Taking a URL apart (reference solution)
package endpoint

import (
	"fmt"
	"net/url"
)

// Endpoint is the part of a URL a client needs to connect.
type Endpoint struct {
	Scheme, Host, Port, Path, Token string
}

var defaultPorts = map[string]string{"http": "80", "https": "443"}

// Parse splits an http(s) URL into an Endpoint.
func Parse(raw string) (Endpoint, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return Endpoint{}, err
	}
	def, ok := defaultPorts[u.Scheme]
	if !ok {
		return Endpoint{}, fmt.Errorf("%q: scheme must be http or https", raw)
	}
	if u.Hostname() == "" {
		return Endpoint{}, fmt.Errorf("%q: no host", raw)
	}
	port := u.Port()
	if port == "" {
		port = def
	}
	return Endpoint{
		Scheme: u.Scheme,
		Host:   u.Hostname(), // u.Host would include the port
		Port:   port,
		Path:   u.Path,
		Token:  u.Query().Get("token"),
	}, nil
}
