// http_client_01: Status codes and closing bodies (reference solution)
package title

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchTitle GETs url and returns the "title" field of the JSON response.
func FetchTitle(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET %s: %s", url, resp.Status)
	}
	var doc struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return "", err
	}
	return doc.Title, nil
}
