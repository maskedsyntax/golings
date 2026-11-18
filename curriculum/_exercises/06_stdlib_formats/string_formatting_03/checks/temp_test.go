package temp

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	for _, tt := range []struct {
		c    Celsius
		want string
	}{
		{21.5, "21.5°C"},
		{-3, "-3.0°C"},
		{0, "0.0°C"},
		{36.66, "36.7°C"},
	} {
		if got := tt.c.String(); got != tt.want {
			t.Errorf("Celsius(%g).String() = %q, want %q", float64(tt.c), got, tt.want)
		}
	}
	if got, want := fmt.Sprint(Reading{"Oslo", -3}), "{Oslo -3.0°C}"; got != want {
		t.Errorf("fmt.Sprint(Reading) = %q, want %q", got, want)
	}
}

func TestReport(t *testing.T) {
	got := Report([]Reading{{"Oslo", -3}, {"Lisbon", 21.5}})
	want := "Oslo   -3.0°C\nLisbon 21.5°C\n"
	if got != want {
		t.Errorf("Report =\n%s\nwant\n%s", got, want)
	}
}
