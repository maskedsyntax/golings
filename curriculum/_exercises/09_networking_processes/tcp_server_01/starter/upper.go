// tcp_server_01: Serving many clients
//
// The server handles one client at a time. Fix Serve. See README.md.
//
// Reference: https://gobyexample.com/tcp-server
package upper

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Serve accepts connections on ln until ln is closed.
func Serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		handle(conn)
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
