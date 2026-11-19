package users

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func api(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			http.Error(w, "want application/json, got "+ct, http.StatusUnsupportedMediaType)
			return
		}
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil || u.Name == "" {
			http.Error(w, "bad user", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"id":"u-%s"}`, u.Name)
	})
	mux.HandleFunc("POST /slow/users", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-time.After(3 * time.Second):
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, `{"id":"too-late"}`)
		}
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func TestCreateUser(t *testing.T) {
	srv := api(t)
	id, err := CreateUser(context.Background(), srv.Client(), srv.URL, User{"ada", "ada@example.com"})
	if err != nil || id != "u-ada" {
		t.Fatalf("CreateUser = %q, %v; want \"u-ada\", nil", id, err)
	}
}

func TestCreateUserRejected(t *testing.T) {
	srv := api(t)
	if id, err := CreateUser(context.Background(), srv.Client(), srv.URL, User{}); err == nil {
		t.Fatalf("CreateUser(empty user) = %q, nil; the server answered 400, so want an error", id)
	}
}

func TestCreateUserDeadline(t *testing.T) {
	srv := api(t)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err := CreateUser(ctx, srv.Client(), srv.URL+"/slow", User{"bob", "bob@example.com"})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("CreateUser error = %v, want one wrapping context.DeadlineExceeded", err)
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("CreateUser took %v; it should give up when ctx expires", d)
	}
}
