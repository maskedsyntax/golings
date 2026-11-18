// Mutant: strings are not quoted.
package describe

import "fmt"

func Describe(x any) string {
	switch v := x.(type) {
	case nil:
		return "nil"
	case int:
		return fmt.Sprintf("int %d", v)
	case string:
		return "string " + v
	case []int:
		return fmt.Sprintf("%d ints", len(v))
	case fmt.Stringer:
		return "stringer " + v.String()
	default:
		return "unknown"
	}
}
