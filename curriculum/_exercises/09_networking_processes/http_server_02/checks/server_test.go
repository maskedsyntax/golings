package server

import (
	"context"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestGracefulShutdown(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	started, release := make(chan struct{}), make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "ok") })
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		close(started)
		select {
		case <-release:
		case <-time.After(5 * time.Second):
		}
		io.WriteString(w, "done")
	})

	ctx, cancel := context.WithCancel(context.Background())
	served := make(chan error, 1)
	go func() { served <- Serve(ctx, ln, mux, 5*time.Second) }()

	base := "http://" + ln.Addr().String()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(base + "/fast")
	if err != nil {
		t.Fatalf("GET /fast: %v", err)
	}
	resp.Body.Close()

	type result struct {
		body string
		err  error
	}
	slow := make(chan result, 1)
	go func() {
		resp, err := client.Get(base + "/slow")
		if err != nil {
			slow <- result{err: err}
			return
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		slow <- result{string(b), err}
	}()

	select {
	case <-started:
	case <-time.After(3 * time.Second):
		t.Fatal("the slow request never reached the handler")
	}
	cancel() // shut down while /slow is in flight
	time.Sleep(100 * time.Millisecond)
	close(release)

	select {
	case r := <-slow:
		if r.err != nil || r.body != "done" {
			t.Errorf("in-flight request got %q, %v; it should have been allowed to finish", r.body, r.err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("the in-flight request never finished")
	}
	select {
	case err := <-served:
		if err != nil {
			t.Errorf("Serve returned %v after a clean shutdown; want nil", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Serve did not return after ctx was cancelled")
	}
	if _, err := net.DialTimeout("tcp", ln.Addr().String(), time.Second); err == nil {
		t.Error("the server still accepts connections after shutting down")
	}
}
