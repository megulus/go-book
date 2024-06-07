package ch8_goroutines

import (
	"io"
	"net"
	"time"
)

/*
TCP server that writes the current time to the client
once per second.
*/

func HandleConn(c net.Conn) {
	defer c.Close() // ignoring errors
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
