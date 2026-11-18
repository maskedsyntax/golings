// Mutant: exported but untagged, so keys are capitalized.
package user

import "encoding/json"

type User struct {
	Name  string
	Email string
	Admin bool `json:",omitempty"`
}

func NewUser(name, email string, admin bool) User {
	return User{name, email, admin}
}

func Encode(u User) (string, error) {
	b, err := json.Marshal(u)
	return string(b), err
}
