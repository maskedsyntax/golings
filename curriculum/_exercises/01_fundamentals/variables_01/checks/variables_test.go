package variables

import "testing"

func TestDescribe(t *testing.T) {
	tests := []struct {
		name        string
		items       []string
		wantCount   int
		wantLongest string
	}{
		{"several", []string{"go", "gopher", "fun"}, 3, "gopher"},
		{"empty", nil, 0, ""},
		{"single", []string{"x"}, 1, "x"},
		{"tie keeps first", []string{"ab", "cd", "e"}, 3, "ab"},
		{"longest last", []string{"a", "bb", "ccc"}, 3, "ccc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count, longest := Describe(tt.items)
			if count != tt.wantCount || longest != tt.wantLongest {
				t.Errorf("Describe(%q) = (%d, %q), want (%d, %q)",
					tt.items, count, longest, tt.wantCount, tt.wantLongest)
			}
		})
	}
}
