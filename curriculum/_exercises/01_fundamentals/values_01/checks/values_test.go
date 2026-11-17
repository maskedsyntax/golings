package values

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		name       string
		sum, count int
		want       float64
	}{
		{"whole", 9, 3, 3},
		{"fraction", 7, 2, 3.5},
		{"negative", -7, 2, -3.5},
		{"zero count", 5, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.sum, tt.count); got != tt.want {
				t.Errorf("Average(%d, %d) = %v, want %v", tt.sum, tt.count, got, tt.want)
			}
		})
	}
}

func TestLabel(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want string
	}{
		{"Gopher", 13, "Gopher is 13"},
		{"Baby", 0, "Baby is 0"},
		{"Ferris", 120, "Ferris is 120"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Label(tt.name, tt.age); got != tt.want {
				t.Errorf("Label(%q, %d) = %q, want %q", tt.name, tt.age, got, tt.want)
			}
		})
	}
}
