// execing_processes_01: Preparing an exec
//
// Implement Prepare. See README.md.
//
// Reference: https://gobyexample.com/execing-processes
package prepare

import "errors"

// Prepare computes the path, argv and environment for executing tool.
func Prepare(tool string, args []string, extra map[string]string, environ []string,
	lookPath func(string) (string, error)) (path string, argv, env []string, err error) {
	// TODO
	return "", nil, nil, errors.New("not implemented")
}
