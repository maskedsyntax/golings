// Mutant: splits at every "=", so "a=b" loses its tail.
package kv

import "strings"

func ParseKV(line string) map[string]string {
	m := make(map[string]string)
	for _, seg := range strings.Split(line, ";") {
		parts := strings.Split(seg, "=")
		k := strings.TrimSpace(parts[0])
		if len(parts) < 2 || k == "" {
			continue
		}
		m[k] = strings.TrimSpace(parts[1])
	}
	return m
}
