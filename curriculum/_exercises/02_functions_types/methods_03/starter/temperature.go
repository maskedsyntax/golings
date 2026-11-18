// methods_03: Methods on named types
//
// Implement Fahrenheit, String and Warmest. See README.md.
//
//	Celsius(100).Fahrenheit() == 212
//	fmt.Sprint(Celsius(21.5)) == "21.5°C"
//	Readings{3, 9, -2}.Warmest() == 9, true
//
// Reference: https://gobyexample.com/methods
package temperature

// Celsius is a temperature in degrees Celsius.
type Celsius float64

// Readings is a series of temperatures.
type Readings []Celsius

// TODO: implement
//   func (c Celsius) Fahrenheit() float64
//   func (c Celsius) String() string
//   func (r Readings) Warmest() (Celsius, bool)
