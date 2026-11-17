// maps_01: Counting and membership
//
// Implement WordCount and Has.
//
//	WordCount([]string{"go", "Go", "go"}) == map[string]int{"go": 2, "Go": 1}
//	Has(map[string]int{"zero": 0}, "zero") == true
//
// Reference: https://gobyexample.com/maps
package maps

// WordCount returns how many times each word occurs (case-sensitive).
// It never returns a nil map.
func WordCount(words []string) map[string]int {
	// TODO: implement WordCount.
	return nil
}

// Has reports whether key is present in m, even when its value is zero.
func Has(m map[string]int, key string) bool {
	// TODO: this is wrong for keys whose value is 0.
	return m[key] != 0
}
