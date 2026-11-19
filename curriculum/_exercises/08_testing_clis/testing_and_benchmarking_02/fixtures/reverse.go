package reverse

// Reverse returns s with its characters (runes) in reverse order.
//
// It is declared as a variable so golings can swap in broken versions and
// check that your tests notice. Call it like any other function.
var Reverse = func(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
