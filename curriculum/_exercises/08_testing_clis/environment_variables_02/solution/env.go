// environment_variables_02: Prefixed settings (reference solution)
package env

import "strings"

// Load returns the variables in environ whose names start with prefix,
// keyed by the lowercased rest of the name.
func Load(environ []string, prefix string) map[string]string {
	out := make(map[string]string)
	for _, kv := range environ {
		key, value, ok := strings.Cut(kv, "=") // values may contain "="
		if !ok {
			continue
		}
		name, ok := strings.CutPrefix(key, prefix)
		if !ok || name == "" {
			continue
		}
		out[strings.ToLower(name)] = value
	}
	return out
}
