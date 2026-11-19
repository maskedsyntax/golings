// Mutant: forgets the Content-Type header.
package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(ctx context.Context, client *http.Client, baseURL string, u User) (string, error) {
	body, _ := json.Marshal(u)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/users", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
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
	err = json.NewDecoder(resp.Body).Decode(&out)
	return out.ID, err
}
