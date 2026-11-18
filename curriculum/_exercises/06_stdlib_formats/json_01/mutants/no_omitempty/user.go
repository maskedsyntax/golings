// Mutant: always writes admin.
package user

import "encoding/json"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Admin bool   `json:"admin"`
}

func NewUser(name, email string, admin bool) User {
	return User{name, email, admin}
}

func Encode(u User) (string, error) {
	b, err := json.Marshal(u)
	return string(b), err
}
