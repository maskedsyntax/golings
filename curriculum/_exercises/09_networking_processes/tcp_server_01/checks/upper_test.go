package upper

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func start(t *testing.T) (string, chan error, net.Listener) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- Serve(ln) }()
	t.Cleanup(func() { ln.Close() })
	return ln.Addr().String(), done, ln
}

func ask(t *testing.T, addr, line string) string {
	t.Helper()
	c, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(2 * time.Second))
	if _, err := c.Write([]byte(line + "\n")); err != nil {
		t.Fatal(err)
	}
	got, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		t.Fatalf("no answer to %q: %v", line, err)
	}
	return got
}

func TestUpper(t *testing.T) {
	addr, _, _ := start(t)
	if got := ask(t, addr, "hello, tcp"); got != "HELLO, TCP\n" {
		t.Fatalf("answer = %q, want %q", got, "HELLO, TCP\n")
	}
}

func TestConcurrentClients(t *testing.T) {
	addr, _, _ := start(t)
	idle, err := net.Dial("tcp", addr) // connects and says nothing
	if err != nil {
		t.Fatal(err)
	}
	defer idle.Close()
	time.Sleep(50 * time.Millisecond)
	if got := ask(t, addr, "still there?"); got != "STILL THERE?\n" {
		t.Fatalf("answer = %q", got)
	}
}

func TestStopsWhenClosed(t *testing.T) {
	_, done, ln := start(t)
	ln.Close()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Serve returned %v after the listener was closed; want nil", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Serve did not return after the listener was closed")
	}
}
