// string_functions_02: Key=value pairs (reference solution)
package kv

import "strings"

// ParseKV parses "k=v; k2=v2" into a map.
func ParseKV(line string) map[string]string {
	m := make(map[string]string)
	for _, seg := range strings.Split(line, ";") {
		k, v, ok := strings.Cut(seg, "=") // splits at the first "=" only
		k = strings.TrimSpace(k)
		if !ok || k == "" {
			continue
		}
		m[k] = strings.TrimSpace(v)
	}
	return m
}
