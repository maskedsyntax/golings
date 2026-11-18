// Mutant: String has a pointer receiver, so fmt doesn't use it for values.
package temperature

import "fmt"

type Celsius float64

type Readings []Celsius

func (c Celsius) Fahrenheit() float64 { return float64(c)*9/5 + 32 }

func (c *Celsius) String() string { return fmt.Sprintf("%.1f°C", float64(*c)) }

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
