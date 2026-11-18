// constants_01: Typed and untyped constants
//
// This file does not compile. Make Factor and TB work without changing
// what they mean.
//
// Reference: https://gobyexample.com/constants
package constants

// Factor scales measurements.
const Factor int = 2

// TB is the number of bytes in a terabyte.
const TB int32 = 1 << 40

// Scale multiplies x by Factor.
func Scale(x float64) float64 {
	return x * Factor
}
