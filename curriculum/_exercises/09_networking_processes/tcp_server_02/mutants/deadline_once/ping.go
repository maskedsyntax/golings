// Mutant: one deadline for the whole connection.
package ping

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Idle time.Duration
}

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
	conn.SetReadDeadline(time.Now().Add(s.Idle))
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
