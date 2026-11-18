package endpoint

import "testing"

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		raw  string
		want Endpoint
	}{
		{"https://api.example.com:8443/v1/items?token=abc&x=1",
			Endpoint{"https", "api.example.com", "8443", "/v1/items", "abc"}},
		{"http://example.com", Endpoint{"http", "example.com", "80", "", ""}},
		{"https://example.com/a%20b?token=x%2By", Endpoint{"https", "example.com", "443", "/a b", "x+y"}},
		{"http://[::1]:9000/health", Endpoint{"http", "::1", "9000", "/health", ""}},
	} {
		got, err := Parse(tt.raw)
		if err != nil {
			t.Errorf("Parse(%q) returned an error: %v", tt.raw, err)
			continue
		}
		if got != tt.want {
			t.Errorf("Parse(%q) =\n  %+v\nwant\n  %+v", tt.raw, got, tt.want)
		}
	}
}

func TestParseErrors(t *testing.T) {
	for _, raw := range []string{"ftp://example.com/file", "https:///no-host", "://bad", "example.com"} {
		if _, err := Parse(raw); err == nil {
			t.Errorf("Parse(%q) returned no error", raw)
		}
	}
}
