package env

import (
	"maps"
	"testing"
)

func TestLoad(t *testing.T) {
	environ := []string{
		"APP_NAME=demo",
		"APP_DB_URL=postgres://u:p@h/db?sslmode=off",
		"OTHER=1",
		"APP_=nothing",
		"=C:=C:\\work",
		"APP_EMPTY=",
		"MALFORMED",
		"APPLE=fruit",
	}
	got := Load(environ, "APP_")
	want := map[string]string{"name": "demo", "db_url": "postgres://u:p@h/db?sslmode=off", "empty": ""}
	if !maps.Equal(got, want) {
		t.Fatalf("Load = %q, want %q", got, want)
	}
}

func TestLoadNone(t *testing.T) {
	got := Load([]string{"HOME=/home/x"}, "APP_")
	if got == nil || len(got) != 0 {
		t.Fatalf("Load = %#v, want an empty map", got)
	}
}
