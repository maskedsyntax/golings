// Mutant: concurrent, but reports the normal stop as an error.
package upper

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		fmt.Fprintln(conn, strings.ToUpper(sc.Text()))
	}
}
