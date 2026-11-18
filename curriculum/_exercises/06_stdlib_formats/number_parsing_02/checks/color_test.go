package color

import "testing"

func TestParseColor(t *testing.T) {
	for _, tt := range []struct {
		in      string
		r, g, b uint8
	}{
		{"#ff8800", 255, 136, 0},
		{"#FF8800", 255, 136, 0},
		{"#000000", 0, 0, 0},
		{"#0a0B0c", 10, 11, 12},
		{"#f80", 255, 136, 0},
		{"#fff", 255, 255, 255},
	} {
		r, g, b, err := ParseColor(tt.in)
		if err != nil {
			t.Errorf("ParseColor(%q) returned an error: %v", tt.in, err)
			continue
		}
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("ParseColor(%q) = %d, %d, %d; want %d, %d, %d", tt.in, r, g, b, tt.r, tt.g, tt.b)
		}
	}
}

func TestParseColorErrors(t *testing.T) {
	for _, in := range []string{
		"ff8800", "#ff880", "#ff88000", "#gg0000", "#-10000", "#+f8800", "#", "", "#ff 800",
	} {
		if _, _, _, err := ParseColor(in); err == nil {
			t.Errorf("ParseColor(%q) returned no error", in)
		}
	}
}
