// http_client_02: POSTing JSON with a deadline
//
// Implement CreateUser. See README.md.
//
// Reference: https://gobyexample.com/http-client
package users

import (
	"context"
	"errors"
	"net/http"
)

// User is the payload for a new account.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser POSTs u to baseURL/users and returns the new user's id.
func CreateUser(ctx context.Context, client *http.Client, baseURL string, u User) (string, error) {
	// TODO
	return "", errors.New("not implemented")
}
