package classify

import "testing"

func TestClassify(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{"42", "positive"}, {" 42 ", "positive"}, {"-7", "negative"}, {"0", "zero"}, {"-0", "zero"},
		{"abc", "not a number"}, {"", "not a number"}, {"4.5", "not a number"},
	} {
		if got := Classify(tt.in); got != tt.want {
			t.Errorf("Classify(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
