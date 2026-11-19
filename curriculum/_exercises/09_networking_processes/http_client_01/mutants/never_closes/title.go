// Mutant: checks the status, never closes.
package title

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchTitle(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
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
