// functions_01: Parameter and result types
//
// This file does not compile. Fix the function declarations.
//
//	Add(2, 3) == 5
//	Greeting("Gopher") == "Hello, Gopher!"
//
// Reference: https://gobyexample.com/functions
package functions

// Add returns the sum of a and b.
func Add(a, b) {
	return a + b
}

// Greeting returns a friendly greeting for name.
func Greeting(name string) string {
	"Hello, " + name + "!"
}
