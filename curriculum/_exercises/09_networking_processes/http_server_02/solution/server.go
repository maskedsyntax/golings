// http_server_02: Graceful shutdown (reference solution)
package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

// Serve serves h on ln until ctx is cancelled, then shuts down gracefully,
// waiting up to grace for in-flight requests.
func Serve(ctx context.Context, ln net.Listener, h http.Handler, grace time.Duration) error {
	srv := &http.Server{Handler: h}
	served := make(chan error, 1)
	go func() { served <- srv.Serve(ln) }()

	select {
	case err := <-served:
		return err // failed before we were asked to stop
	case <-ctx.Done():
	}

	sctx, cancel := context.WithTimeout(context.Background(), grace)
	defer cancel()
	if err := srv.Shutdown(sctx); err != nil {
		return err
	}
	if err := <-served; !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
