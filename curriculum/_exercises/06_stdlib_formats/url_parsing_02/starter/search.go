// url_parsing_02: Building query strings
//
// SearchURL doesn't escape anything. Fix it. See README.md.
//
// Reference: https://gobyexample.com/url-parsing
package search

import "strconv"

// SearchURL returns base with the query parameters q and page set.
func SearchURL(base, q string, page int) (string, error) {
	return base + "?q=" + q + "&page=" + strconv.Itoa(page), nil
}
