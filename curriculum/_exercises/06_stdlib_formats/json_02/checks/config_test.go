package config

import "testing"

func TestParseConfig(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want Config
	}{
		{`{}`, Config{Host: "localhost", Port: 8080}},
		{`{"port": 9000}`, Config{Host: "localhost", Port: 9000}},
		{`{"host": "example.com", "port": 443, "debug": true}`, Config{Host: "example.com", Port: 443, Debug: true}},
	} {
		got, err := ParseConfig([]byte(tt.in))
		if err != nil {
			t.Errorf("ParseConfig(%s) returned an error: %v", tt.in, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseConfig(%s) = %+v, want %+v", tt.in, got, tt.want)
		}
	}
}

func TestParseConfigErrors(t *testing.T) {
	for _, in := range []string{
		`{"prot": 9000}`,
		`{"port": "80"}`,
		`{"port": 0}`,
		`{"port": 70000}`,
		`not json`,
		``,
	} {
		if _, err := ParseConfig([]byte(in)); err == nil {
			t.Errorf("ParseConfig(%q) returned no error", in)
		}
	}
}
