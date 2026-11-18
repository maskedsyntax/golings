package temperature

import (
	"fmt"
	"math"
	"testing"
)

func TestFahrenheit(t *testing.T) {
	for _, tt := range []struct {
		c    Celsius
		want float64
	}{{0, 32}, {100, 212}, {-40, -40}, {37, 98.6}} {
		if got := tt.c.Fahrenheit(); math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("Celsius(%v).Fahrenheit() = %v, want %v", float64(tt.c), got, tt.want)
		}
	}
}

func TestString(t *testing.T) {
	for _, tt := range []struct {
		c    Celsius
		want string
	}{{21.5, "21.5°C"}, {-3, "-3.0°C"}, {0, "0.0°C"}, {36.66, "36.7°C"}} {
		if got := tt.c.String(); got != tt.want {
			t.Errorf("Celsius(%v).String() = %q, want %q", float64(tt.c), got, tt.want)
		}
		if got := fmt.Sprint(tt.c); got != tt.want {
			t.Errorf("fmt.Sprint(Celsius(%v)) = %q, want %q (does fmt see your String method?)", float64(tt.c), got, tt.want)
		}
	}
	if got := fmt.Sprintf("now %v", Celsius(5)); got != "now 5.0°C" {
		t.Errorf("Sprintf with %%v = %q, want %q", got, "now 5.0°C")
	}
}

func TestWarmest(t *testing.T) {
	if got, ok := (Readings{3, 9, -2}).Warmest(); !ok || got != 9 {
		t.Errorf("Readings{3, 9, -2}.Warmest() = %v, %v, want 9, true", float64(got), ok)
	}
	if got, ok := (Readings{-5, -1, -9}).Warmest(); !ok || got != -1 {
		t.Errorf("Readings{-5, -1, -9}.Warmest() = %v, %v, want -1, true", float64(got), ok)
	}
	var none Readings
	if got, ok := none.Warmest(); ok || got != 0 {
		t.Errorf("empty Readings.Warmest() = %v, %v, want 0, false", float64(got), ok)
	}
}
