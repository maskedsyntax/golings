// tcp_server_01: Serving many clients (reference solution)
package upper

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

// Serve accepts connections on ln until ln is closed.
func Serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if errors.Is(err, net.ErrClosed) {
			return nil // the listener was closed: a normal stop
		}
		if err != nil {
			return err
		}
		go handle(conn) // one goroutine per client
	}
}

// handle uppercases every line the client sends.
func handle(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		fmt.Fprintln(conn, strings.ToUpper(sc.Text()))
	}
}
