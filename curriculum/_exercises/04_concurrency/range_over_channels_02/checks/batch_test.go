package batch

import (
	"reflect"
	"testing"
	"time"
)

func feed(items ...string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, s := range items {
			ch <- s
		}
	}()
	return ch
}

func collect(t *testing.T, ch <-chan []string) [][]string {
	t.Helper()
	if ch == nil {
		t.Fatal("Batch returned a nil channel")
	}
	var out [][]string
	timeout := time.After(2 * time.Second)
	for {
		select {
		case b, ok := <-ch:
			if !ok {
				return out
			}
			out = append(out, b)
		case <-timeout:
			t.Fatalf("got batches %q and then the output was never closed", out)
		}
	}
}

func TestBatch(t *testing.T) {
	for _, tt := range []struct {
		items []string
		size  int
		want  [][]string
	}{
		{[]string{"a", "b", "c", "d", "e"}, 2, [][]string{{"a", "b"}, {"c", "d"}, {"e"}}},
		{[]string{"a", "b", "c", "d"}, 2, [][]string{{"a", "b"}, {"c", "d"}}},
		{[]string{"a", "b"}, 5, [][]string{{"a", "b"}}},
		{nil, 3, nil},
	} {
		if got := collect(t, Batch(feed(tt.items...), tt.size)); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Batch(%q, %d) = %q, want %q", tt.items, tt.size, got, tt.want)
		}
	}
}
