// Mutant: graceful, but returns http.ErrServerClosed.
package server

import (
	"context"
	"net"
	"net/http"
	"time"
)

func Serve(ctx context.Context, ln net.Listener, h http.Handler, grace time.Duration) error {
	srv := &http.Server{Handler: h}
	go func() {
		<-ctx.Done()
		sctx, cancel := context.WithTimeout(context.Background(), grace)
		defer cancel()
		srv.Shutdown(sctx)
	}()
	return srv.Serve(ln)
}
