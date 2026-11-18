package functions

import "testing"

func TestAdd(t *testing.T) {
	for _, tt := range []struct{ a, b, want int }{{2, 3, 5}, {-4, 4, 0}, {0, 0, 0}, {100, -1, 99}} {
		if got := Add(tt.a, tt.b); got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestGreeting(t *testing.T) {
	for _, tt := range []struct{ name, want string }{{"Gopher", "Hello, Gopher!"}, {"", "Hello, !"}} {
		if got := Greeting(tt.name); got != tt.want {
			t.Errorf("Greeting(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}
