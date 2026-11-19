// http_client_02: POSTing JSON with a deadline (reference solution)
package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// User is the payload for a new account.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser POSTs u to baseURL/users and returns the new user's id.
func CreateUser(ctx context.Context, client *http.Client, baseURL string, u User) (string, error) {
	body, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/users", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req) // fails with ctx's error when ctx expires
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("create user: %s", resp.Status)
	}
	var out struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.ID, nil
}
