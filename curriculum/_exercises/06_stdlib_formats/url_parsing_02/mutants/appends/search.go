// Mutant: Add instead of Set, so an existing page is duplicated.
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
	values := u.Query()
	values.Add("q", q)
	values.Add("page", strconv.Itoa(page))
	u.RawQuery = values.Encode()
	return u.String(), nil
}
