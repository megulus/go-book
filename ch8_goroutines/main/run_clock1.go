package main

import (
	ch8 "github.com/megulus/go-book/ch8_goroutines"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		ch8.HandleConn(conn) // handle one connection at a time
	}
}
