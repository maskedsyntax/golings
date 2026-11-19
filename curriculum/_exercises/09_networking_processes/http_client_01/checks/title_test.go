package title

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// trackingTransport counts response bodies that were never closed.
type trackingTransport struct {
	next http.RoundTripper
	open atomic.Int32
}

type trackedBody struct {
	io.ReadCloser
	t      *trackingTransport
	closed atomic.Bool
}

func (b *trackedBody) Close() error {
	if b.closed.CompareAndSwap(false, true) {
		b.t.open.Add(-1)
	}
	return b.ReadCloser.Close()
}

func (t *trackingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := t.next.RoundTrip(r)
	if err == nil {
		t.open.Add(1)
		resp.Body = &trackedBody{ReadCloser: resp.Body, t: t}
	}
	return resp, err
}

func server(t *testing.T) (*httptest.Server, *http.Client, *trackingTransport) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"title":"Go by Example"}`)
	})
	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"error":"no such page"}`)
	})
	mux.HandleFunc("/broken", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{}`)
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	tr := &trackingTransport{next: srv.Client().Transport}
	return srv, &http.Client{Transport: tr}, tr
}

func TestFetchTitle(t *testing.T) {
	srv, client, _ := server(t)
	got, err := FetchTitle(client, srv.URL+"/ok")
	if err != nil || got != "Go by Example" {
		t.Fatalf("FetchTitle(/ok) = %q, %v; want \"Go by Example\", nil", got, err)
	}
}

func TestFetchTitleStatus(t *testing.T) {
	srv, client, _ := server(t)
	for path, code := range map[string]string{"/missing": "404", "/broken": "500"} {
		got, err := FetchTitle(client, srv.URL+path)
		if err == nil {
			t.Errorf("FetchTitle(%s) = %q, nil; want an error", path, got)
		} else if !strings.Contains(err.Error(), code) {
			t.Errorf("FetchTitle(%s) error %q should mention the status %s", path, err, code)
		}
	}
}

func TestFetchTitleClosesBodies(t *testing.T) {
	srv, client, tr := server(t)
	for _, path := range []string{"/ok", "/missing", "/broken", "/ok"} {
		FetchTitle(client, srv.URL+path)
	}
	if n := tr.open.Load(); n != 0 {
		t.Fatalf("%d response bodies were never closed", n)
	}
}
