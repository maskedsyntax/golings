// Mutant: strings are not quoted, so "" is indistinguishable.
package describe

import "fmt"

func Describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case int:
		return fmt.Sprintf("int %d", x)
	case string:
		return "string " + x
	case bool:
		return fmt.Sprintf("bool %t", x)
	case []int:
		return fmt.Sprintf("[]int of length %d", len(x))
	default:
		return "unknown"
	}
}
