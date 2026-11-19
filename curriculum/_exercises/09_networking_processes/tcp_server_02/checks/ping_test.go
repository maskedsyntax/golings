package ping

import (
	"bufio"
	"errors"
	"net"
	"testing"
	"time"
)

func start(t *testing.T, idle time.Duration) string {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	s := &Server{Idle: idle}
	go s.Serve(ln)
	t.Cleanup(func() { ln.Close() })
	return ln.Addr().String()
}

func dial(t *testing.T, addr string) net.Conn {
	t.Helper()
	c, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Close() })
	return c
}

func expectHangUp(t *testing.T, c net.Conn, r *bufio.Reader) {
	t.Helper()
	c.SetReadDeadline(time.Now().Add(3 * time.Second))
	_, err := r.ReadString('\n')
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		t.Fatal("the server never closed the idle connection")
	}
	if err == nil {
		t.Fatal("the server sent data instead of closing the idle connection")
	}
}

func TestPing(t *testing.T) {
	c := dial(t, start(t, time.Second))
	r := bufio.NewReader(c)
	c.SetDeadline(time.Now().Add(2 * time.Second))
	for _, tt := range []struct{ send, want string }{{"PING", "PONG\n"}, {"HELLO", "ERR unknown command\n"}} {
		c.Write([]byte(tt.send + "\n"))
		if got, err := r.ReadString('\n'); err != nil || got != tt.want {
			t.Fatalf("sent %q, got %q, %v; want %q", tt.send, got, err, tt.want)
		}
	}
}

func TestIdleClientIsDropped(t *testing.T) {
	c := dial(t, start(t, 200*time.Millisecond))
	expectHangUp(t, c, bufio.NewReader(c))
}

func TestActiveClientStays(t *testing.T) {
	c := dial(t, start(t, 400*time.Millisecond))
	r := bufio.NewReader(c)
	end := time.Now().Add(1200 * time.Millisecond) // three times the idle limit
	for time.Now().Before(end) {
		c.SetDeadline(time.Now().Add(2 * time.Second))
		if _, err := c.Write([]byte("PING\n")); err != nil {
			t.Fatalf("the server hung up on an active client: %v", err)
		}
		if got, err := r.ReadString('\n'); err != nil || got != "PONG\n" {
			t.Fatalf("the server hung up on an active client (got %q, %v)", got, err)
		}
		time.Sleep(50 * time.Millisecond)
	}
	expectHangUp(t, c, r) // and once it goes quiet, it is dropped
}
