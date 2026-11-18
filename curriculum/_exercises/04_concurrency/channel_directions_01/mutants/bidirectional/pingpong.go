// Mutant: compiles by dropping the directions entirely.
package pingpong

func Ping(pings chan string, msg string) { pings <- msg }

func Pong(pings chan string, pongs chan string) { pongs <- <-pings }
