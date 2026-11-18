package search

import (
	"net/url"
	"testing"
)

func check(t *testing.T, base, q string, page int, wantPath string, want url.Values) {
	t.Helper()
	got, err := SearchURL(base, q, page)
	if err != nil {
		t.Fatalf("SearchURL(%q, %q, %d) returned an error: %v", base, q, page, err)
	}
	u, err := url.Parse(got)
	if err != nil {
		t.Fatalf("SearchURL returned %q, which doesn't parse: %v", got, err)
	}
	if u.Path != wantPath {
		t.Errorf("SearchURL returned %q: path %q, want %q", got, u.Path, wantPath)
	}
	gotQ := u.Query()
	if len(gotQ) != len(want) {
		t.Errorf("SearchURL returned %q: parameters %v, want %v", got, gotQ, want)
	}
	for k, v := range want {
		if gotQ.Get(k) != v[0] || len(gotQ[k]) != 1 {
			t.Errorf("SearchURL returned %q: %s = %q, want exactly %q", got, k, gotQ[k], v[0])
		}
	}
}

func TestSearchURL(t *testing.T) {
	check(t, "https://example.com/search", "golang", 1, "/search",
		url.Values{"q": {"golang"}, "page": {"1"}})
	check(t, "https://example.com/search", "go & rust", 2, "/search",
		url.Values{"q": {"go & rust"}, "page": {"2"}})
	check(t, "https://example.com/search", "a+b=100%", 3, "/search",
		url.Values{"q": {"a+b=100%"}, "page": {"3"}})
	check(t, "https://example.com/search?lang=en", "#tags", 4, "/search",
		url.Values{"q": {"#tags"}, "page": {"4"}, "lang": {"en"}})
	check(t, "https://example.com/search?page=9", "x", 5, "/search",
		url.Values{"q": {"x"}, "page": {"5"}})
}

func TestSearchURLBadBase(t *testing.T) {
	if _, err := SearchURL("://nope", "x", 1); err == nil {
		t.Fatal("SearchURL with an unparsable base returned no error")
	}
}
