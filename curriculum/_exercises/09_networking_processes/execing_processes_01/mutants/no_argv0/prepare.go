// Mutant: argv without the program name.
package prepare

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func Prepare(tool string, args []string, extra map[string]string, environ []string,
	lookPath func(string) (string, error)) (string, []string, []string, error) {
	path, err := lookPath(tool)
	if err != nil {
		return "", nil, nil, fmt.Errorf("finding %s: %w", tool, err)
	}
	var env []string
	used := map[string]bool{}
	for _, kv := range environ {
		key, _, _ := strings.Cut(kv, "=")
		if v, ok := extra[key]; ok {
			kv = key + "=" + v
			used[key] = true
		}
		env = append(env, kv)
	}
	for _, key := range slices.Sorted(maps.Keys(extra)) {
		if !used[key] {
			env = append(env, key+"="+extra[key])
		}
	}
	return path, args, env, nil
}
