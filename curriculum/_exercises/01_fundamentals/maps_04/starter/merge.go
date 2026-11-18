// maps_04: Writing to a nil map
//
// Merge panics. Fix it without modifying the input maps.
//
//	Merge({"x": 1, "y": 2}, {"y": 3}) == {"x": 1, "y": 5}
//
// Reference: https://gobyexample.com/maps
package merge

// Merge returns a new map containing the sum of a and b per key.
func Merge(a, b map[string]int) map[string]int {
	var out map[string]int
	for k, v := range a {
		out[k] += v
	}
	for k, v := range b {
		out[k] += v
	}
	return out
}
