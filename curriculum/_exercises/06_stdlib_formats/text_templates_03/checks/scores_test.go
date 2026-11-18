package scores

import "testing"

func TestRender(t *testing.T) {
	got, err := Render([]Score{{"Ada", 9}, {"Bob", 7}})
	if err != nil {
		t.Fatalf("Render returned an error: %v", err)
	}
	if want := "ADA: 9\nBOB: 7\n"; got != want {
		t.Fatalf("Render = %q, want %q", got, want)
	}
	got, err = Render(nil)
	if err != nil {
		t.Fatalf("Render(nil) returned an error: %v", err)
	}
	if got != "" {
		t.Fatalf("Render(nil) = %q, want \"\"", got)
	}
}
