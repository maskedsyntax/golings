// http_client_01: Status codes and closing bodies
//
// FetchTitle ignores the status code and leaks bodies. Fix it. See README.md.
//
// Reference: https://gobyexample.com/http-client
package title

import (
	"encoding/json"
	"net/http"
)

// FetchTitle GETs url and returns the "title" field of the JSON response.
func FetchTitle(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	var doc struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return "", err
	}
	return doc.Title, nil
}
