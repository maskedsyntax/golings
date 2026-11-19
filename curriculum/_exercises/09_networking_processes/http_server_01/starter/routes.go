// http_server_01: Routing with ServeMux
//
// Implement NewMux. See README.md.
//
// Reference: https://gobyexample.com/http-server
package routes

import "net/http"

// NewMux returns the application's router.
func NewMux() http.Handler {
	mux := http.NewServeMux()
	// TODO: register the routes.
	return mux
}
