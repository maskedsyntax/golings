// Mutant: uses u.Host, which includes the port.
package endpoint

import (
	"fmt"
	"net/url"
)

type Endpoint struct {
	Scheme, Host, Port, Path, Token string
}

func Parse(raw string) (Endpoint, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return Endpoint{}, err
	}
	if (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return Endpoint{}, fmt.Errorf("bad url %q", raw)
	}
	port := u.Port()
	if port == "" {
		port = map[string]string{"http": "80", "https": "443"}[u.Scheme]
	}
	return Endpoint{u.Scheme, u.Host, port, u.Path, u.Query().Get("token")}, nil
}
