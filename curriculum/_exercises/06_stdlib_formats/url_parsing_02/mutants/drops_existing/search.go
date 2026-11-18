// Mutant: escapes correctly but throws away the base's parameters.
package search

import (
	"net/url"
	"strconv"
)

func SearchURL(base, q string, page int) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	u.RawQuery = url.Values{"q": {q}, "page": {strconv.Itoa(page)}}.Encode()
	return u.String(), nil
}
