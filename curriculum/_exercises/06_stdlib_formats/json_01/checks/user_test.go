package user

import "testing"

func TestEncode(t *testing.T) {
	for _, tt := range []struct {
		u    User
		want string
	}{
		{NewUser("Ada", "ada@example.com", true), `{"name":"Ada","email":"ada@example.com","admin":true}`},
		{NewUser("Bob", "bob@example.com", false), `{"name":"Bob","email":"bob@example.com"}`},
	} {
		got, err := Encode(tt.u)
		if err != nil {
			t.Fatalf("Encode returned an error: %v", err)
		}
		if got != tt.want {
			t.Errorf("Encode = %s, want %s", got, tt.want)
		}
	}
}
