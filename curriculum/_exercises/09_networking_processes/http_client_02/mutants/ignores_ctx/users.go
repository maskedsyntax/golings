// Mutant: uses client.Post, which ignores ctx.
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
	resp, err := client.Post(baseURL+"/users", "application/json", bytes.NewReader(body))
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
