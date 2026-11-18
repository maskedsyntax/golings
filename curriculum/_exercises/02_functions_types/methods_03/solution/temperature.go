// methods_03: Methods on named types (reference solution)
package temperature

import "fmt"

// Celsius is a temperature in degrees Celsius.
type Celsius float64

// Readings is a series of temperatures.
type Readings []Celsius

// Fahrenheit converts c to degrees Fahrenheit.
func (c Celsius) Fahrenheit() float64 {
	return float64(c)*9/5 + 32
}

// String formats c with one decimal place, e.g. "21.5°C".
func (c Celsius) String() string {
	// float64(c) matters: formatting c itself with %v would call String again.
	return fmt.Sprintf("%.1f°C", float64(c))
}

// Warmest returns the highest reading. ok is false if there are none.
func (r Readings) Warmest() (Celsius, bool) {
	if len(r) == 0 {
		return 0, false
	}
	best := r[0]
	for _, c := range r[1:] {
		best = max(best, c)
	}
	return best, true
}
