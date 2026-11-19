package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func do(method, target, body string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	NewMux().ServeHTTP(rec, httptest.NewRequest(method, target, strings.NewReader(body)))
	return rec
}

func TestRoutes(t *testing.T) {
	for _, tt := range []struct {
		method, target, body string
		code                 int
		want                 string
	}{
		{"GET", "/hello/Ada", "", 200, "Hello, Ada!\n"},
		{"GET", "/hello/Grace%20Hopper", "", 200, "Hello, Grace Hopper!\n"},
		{"POST", "/echo", "ping", 200, "ping"},
		{"POST", "/echo", "", 200, ""},
	} {
		rec := do(tt.method, tt.target, tt.body)
		if rec.Code != tt.code || rec.Body.String() != tt.want {
			t.Errorf("%s %s: %d %q; want %d %q", tt.method, tt.target, rec.Code, rec.Body.String(), tt.code, tt.want)
		}
	}
}

func TestRouteErrors(t *testing.T) {
	for _, tt := range []struct {
		method, target string
		code           int
	}{
		{"GET", "/echo", http.StatusMethodNotAllowed},
		{"DELETE", "/hello/Ada", http.StatusMethodNotAllowed},
		{"GET", "/hello/", http.StatusNotFound},
		{"GET", "/nope", http.StatusNotFound},
		{"HEAD", "/hello/Ada", http.StatusOK},
	} {
		if rec := do(tt.method, tt.target, ""); rec.Code != tt.code {
			t.Errorf("%s %s: status %d, want %d", tt.method, tt.target, rec.Code, tt.code)
		}
	}
}
