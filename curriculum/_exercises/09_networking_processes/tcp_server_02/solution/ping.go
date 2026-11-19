// tcp_server_02: Idle timeouts (reference solution)
package ping

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

// Server answers PING with PONG and drops clients that are idle for too long.
type Server struct {
	Idle time.Duration
}

// Serve accepts connections on ln until ln is closed.
func (s *Server) Serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if errors.Is(err, net.ErrClosed) {
			return nil
		}
		if err != nil {
			return err
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		// Renew the deadline before every read: it limits each silence, not the session.
		conn.SetReadDeadline(time.Now().Add(s.Idle))
		line, err := r.ReadString('\n')
		if err != nil {
			return // includes the idle timeout
		}
		if line == "PING\n" {
			fmt.Fprint(conn, "PONG\n")
		} else {
			fmt.Fprint(conn, "ERR unknown command\n")
		}
	}
}
