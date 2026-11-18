// json_01: Exported fields and tags
//
// Encode produces {}. Fix the User type. See README.md.
//
// Reference: https://gobyexample.com/json
package user

import "encoding/json"

// User is an account.
type User struct {
	name  string
	email string
	admin bool
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
