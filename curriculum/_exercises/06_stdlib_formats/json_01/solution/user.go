// json_01: Exported fields and tags (reference solution)
package user

import "encoding/json"

// User is an account.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Admin bool   `json:"admin,omitempty"`
}

// NewUser creates a User.
func NewUser(name, email string, admin bool) User {
	return User{name, email, admin}
}

// Encode returns u as JSON.
func Encode(u User) (string, error) {
	b, err := json.Marshal(u)
	return string(b), err
}
