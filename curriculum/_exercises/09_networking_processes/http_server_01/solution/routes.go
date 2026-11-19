// http_server_01: Routing with ServeMux (reference solution)
package routes

import (
	"fmt"
	"io"
	"net/http"
)

// NewMux returns the application's router.
func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", r.PathValue("name"))
	})
	mux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, r.Body)
	})
	return mux
}
