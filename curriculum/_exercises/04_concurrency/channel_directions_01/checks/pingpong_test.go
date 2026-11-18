package pingpong

import (
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	Ping(pings, "passed message")
	Pong(pings, pongs)
	select {
	case got := <-pongs:
		if got != "passed message" {
			t.Fatalf("received %q, want %q", got, "passed message")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("nothing arrived on pongs")
	}
}
