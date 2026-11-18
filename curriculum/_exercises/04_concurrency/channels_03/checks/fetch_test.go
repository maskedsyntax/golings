package fetch

import (
	"slices"
	"strings"
	"testing"
)

// reverseFinisher returns a fetch func whose calls finish in reverse input
// order: the fetch for urls[i] waits until urls[i+1] has finished.
func reverseFinisher(urls []string) func(string) string {
	gates := map[string]chan struct{}{}
	for _, u := range urls {
		gates[u] = make(chan struct{})
	}
	next := map[string]string{}
	for i := 0; i+1 < len(urls); i++ {
		next[urls[i]] = urls[i+1]
	}
	return func(u string) string {
		if n, ok := next[u]; ok {
			<-gates[n]
		}
		defer close(gates[u])
		return strings.ToUpper(u)
	}
}

func TestFetchAllKeepsOrder(t *testing.T) {
	urls := []string{"a.example", "b.example", "c.example", "d.example"}
	got := FetchAll(urls, reverseFinisher(urls))
	want := []string{"A.EXAMPLE", "B.EXAMPLE", "C.EXAMPLE", "D.EXAMPLE"}
	if !slices.Equal(got, want) {
		t.Fatalf("FetchAll = %q\nwant        %q\n(the last fetch finished first)", got, want)
	}
}

func TestFetchAllEmpty(t *testing.T) {
	if got := FetchAll(nil, strings.ToUpper); len(got) != 0 {
		t.Fatalf("FetchAll(nil) = %q", got)
	}
}
