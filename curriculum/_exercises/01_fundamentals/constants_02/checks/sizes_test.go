package sizes

import "testing"

func TestHumanSize(t *testing.T) {
	tests := []struct {
		n    int64
		want string
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1023, "1023 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1048575, "1024.0 KB"},
		{1048576, "1.0 MB"},
		{5 * 1024 * 1024 * 1024, "5.0 GB"},
		{2 * 1024 * 1024 * 1024 * 1024, "2048.0 GB"},
	}
	for _, tt := range tests {
		if got := HumanSize(tt.n); got != tt.want {
			t.Errorf("HumanSize(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
