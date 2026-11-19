// execing_processes_01: Preparing an exec (reference solution)
package prepare

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

// Prepare computes the path, argv and environment for executing tool.
func Prepare(tool string, args []string, extra map[string]string, environ []string,
	lookPath func(string) (string, error)) (path string, argv, env []string, err error) {
	path, err = lookPath(tool)
	if err != nil {
		return "", nil, nil, fmt.Errorf("finding %s: %w", tool, err)
	}
	argv = append([]string{tool}, args...) // argv[0] is the program name

	used := make(map[string]bool)
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
	return path, argv, env, nil
}
