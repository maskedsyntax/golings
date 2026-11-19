// Mutant: splits at every "=".
package env

import "strings"

func Load(environ []string, prefix string) map[string]string {
	out := make(map[string]string)
	for _, kv := range environ {
		parts := strings.Split(kv, "=")
		if len(parts) < 2 {
			continue
		}
		name, ok := strings.CutPrefix(parts[0], prefix)
		if !ok || name == "" {
			continue
		}
		out[strings.ToLower(name)] = parts[1]
	}
	return out
}
