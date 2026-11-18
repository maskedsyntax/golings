package list

import "testing"

func TestRender(t *testing.T) {
	for _, tt := range []struct {
		items []string
		want  string
	}{
		{[]string{"apple", "bread"}, "Shopping list:\n- apple\n- bread\n"},
		{[]string{"milk"}, "Shopping list:\n- milk\n"},
		{nil, "Nothing to buy.\n"},
		{[]string{}, "Nothing to buy.\n"},
	} {
		got, err := Render(tt.items)
		if err != nil {
			t.Fatalf("Render(%q) returned an error: %v", tt.items, err)
		}
		if got != tt.want {
			t.Errorf("Render(%q) = %q, want %q", tt.items, got, tt.want)
		}
	}
}
