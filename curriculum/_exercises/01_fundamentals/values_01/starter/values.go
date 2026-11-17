// values_01: Mixing numbers and strings
//
// This file does not compile. Fix the type errors so that:
//
//	Average(7, 2)     == 3.5
//	Average(5, 0)     == 0
//	Label("Gopher", 13) == "Gopher is 13"
//
// Reference: https://gobyexample.com/values
package values

// Average returns sum/count as an exact (non-truncated) value.
// A count of zero yields 0.
func Average(sum int, count int) float64 {
	if count == 0 {
		return 0
	}
	return sum / count
}

// Label describes a gopher's age, for example "Gopher is 13".
func Label(name string, age int) string {
	return name + " is " + age
}
