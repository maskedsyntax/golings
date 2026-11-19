// Mutant: appends overrides instead of replacing them.
package prepare

import (
	"fmt"
	"maps"
	"slices"
)

func Prepare(tool string, args []string, extra map[string]string, environ []string,
	lookPath func(string) (string, error)) (string, []string, []string, error) {
	path, err := lookPath(tool)
	if err != nil {
		return "", nil, nil, fmt.Errorf("finding %s: %w", tool, err)
	}
	env := slices.Clone(environ)
	for _, key := range slices.Sorted(maps.Keys(extra)) {
		env = append(env, key+"="+extra[key])
	}
	return path, append([]string{tool}, args...), env, nil
}
