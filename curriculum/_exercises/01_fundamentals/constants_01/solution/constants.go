// constants_01: Typed and untyped constants (reference solution)
package constants

// Factor scales measurements.
const Factor = 2

// TB is the number of bytes in a terabyte.
const TB = 1 << 40

// Scale multiplies x by Factor.
func Scale(x float64) float64 {
	return x * Factor
}
