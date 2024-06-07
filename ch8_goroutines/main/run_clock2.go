package main

import (
	"flag"
	"fmt"
	ch8 "github.com/megulus/go-book/ch8_goroutines"
	"log"
	"net"
)

/*
Modifies clock server to accept a port number

to invoke, e.g.:
TZ="UTC" go run ch8_goroutines/main/run_clock2.go --port 8020 &
*/

func main() {
	// p := os.Args[1]
	port := flag.String("port", "8000", "port to listen on")
	flag.Parse()
	fmt.Println("port", *port)
	address := fmt.Sprintf("localhost:%s", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		ch8.HandleConn(conn)
	}
}
