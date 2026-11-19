// http_server_02: Graceful shutdown
//
// Serve drops in-flight requests and reports an error on a normal shutdown.
// Fix it. See README.md.
//
// Reference: https://gobyexample.com/http-server
package server

import (
	"context"
	"net"
	"net/http"
	"time"
)

// Serve serves h on ln until ctx is cancelled, then shuts down gracefully,
// waiting up to grace for in-flight requests.
func Serve(ctx context.Context, ln net.Listener, h http.Handler, grace time.Duration) error {
	srv := &http.Server{Handler: h}
	go func() {
		<-ctx.Done()
		srv.Close()
	}()
	return srv.Serve(ln)
}
