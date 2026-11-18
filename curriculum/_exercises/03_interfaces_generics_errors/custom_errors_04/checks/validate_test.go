package validate

import (
	"errors"
	"testing"
)

// fields returns the Field of every *FieldError joined into err.
func fields(err error) []string {
	var out []string
	multi, ok := err.(interface{ Unwrap() []error })
	if !ok {
		var fe *FieldError
		if errors.As(err, &fe) {
			out = append(out, fe.Field)
		}
		return out
	}
	for _, e := range multi.Unwrap() {
		out = append(out, fields(e)...)
	}
	return out
}

func TestValid(t *testing.T) {
	if err := ValidateAll(User{Name: "Ana", Email: "a@b", Age: 3}); err != nil {
		t.Errorf("ValidateAll(valid) = %#v, want nil", err)
	}
}

func TestAllProblems(t *testing.T) {
	err := ValidateAll(User{Email: "x", Age: -1})
	if err == nil {
		t.Fatal("ValidateAll returned nil for an invalid user")
	}
	want := "name: must not be empty\nemail: must contain @\nage: must not be negative"
	if err.Error() != want {
		t.Errorf("Error() = %q, want %q", err.Error(), want)
	}
	if got := fields(err); len(got) != 3 || got[0] != "name" || got[1] != "email" || got[2] != "age" {
		t.Errorf("joined FieldErrors have fields %q, want [name email age]", got)
	}
	var fe *FieldError
	if !errors.As(err, &fe) || fe.Field != "name" {
		t.Errorf("errors.As should find the first FieldError (name), got %+v", fe)
	}
}

func TestOneProblem(t *testing.T) {
	err := ValidateAll(User{Name: "Ana", Email: "a@b", Age: -5})
	if got := fields(err); len(got) != 1 || got[0] != "age" {
		t.Errorf("fields = %q, want [age]", got)
	}
	if err == nil || err.Error() != "age: must not be negative" {
		t.Errorf("Error() = %v", err)
	}
}
