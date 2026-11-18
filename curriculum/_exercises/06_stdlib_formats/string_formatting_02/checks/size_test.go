package size

import "testing"

func TestFormatBytes(t *testing.T) {
	for _, tt := range []struct {
		n    int64
		want string
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1023, "1023 B"},
		{1024, "1.0 KiB"},
		{1536, "1.5 KiB"},
		{10 << 20, "10.0 MiB"},
		{1 << 30, "1.0 GiB"},
		{5 << 40, "5.0 TiB"},
		{1 << 50, "1024.0 TiB"},
	} {
		if got := FormatBytes(tt.n); got != tt.want {
			t.Errorf("FormatBytes(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
