// url_parsing_02: Building query strings (reference solution)
package search

import (
	"net/url"
	"strconv"
)

// SearchURL returns base with the query parameters q and page set.
func SearchURL(base, q string, page int) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	values := u.Query() // keeps existing parameters
	values.Set("q", q)
	values.Set("page", strconv.Itoa(page))
	u.RawQuery = values.Encode() // escapes every key and value
	return u.String(), nil
}
