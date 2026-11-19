// tcp_server_02: Idle timeouts
//
// Silent clients are never disconnected. Fix handle. See README.md.
//
// Reference: https://gobyexample.com/tcp-server
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
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}
		if line == "PING\n" {
			fmt.Fprint(conn, "PONG\n")
		} else {
			fmt.Fprint(conn, "ERR unknown command\n")
		}
	}
}
