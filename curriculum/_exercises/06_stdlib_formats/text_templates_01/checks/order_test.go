package order

import "testing"

func TestSummary(t *testing.T) {
	got, err := Summary(Order{Name: "Ada", Items: []string{"tea", "cake"}})
	if err != nil {
		t.Fatalf("Summary returned an error: %v", err)
	}
	if want := "Hello, Ada! You ordered 2 items."; got != want {
		t.Fatalf("Summary = %q, want %q", got, want)
	}
	got, err = Summary(Order{Name: "Bob"})
	if err != nil {
		t.Fatalf("Summary returned an error: %v", err)
	}
	if want := "Hello, Bob! You ordered 0 items."; got != want {
		t.Fatalf("Summary = %q, want %q", got, want)
	}
}
