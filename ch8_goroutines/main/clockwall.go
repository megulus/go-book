package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

/*
Write a program that acts as a client of several clock servers at
once, reading the times from each one and displaying the results in
a table.
*/

func ClockWall(city string, port string, ch chan string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Print(err)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Print(err)
			return
		}
		//fmt.Printf("%s: %s\n", city, buffer[:n])
		ch <- fmt.Sprintf("%s: %s", city, buffer[:n])
	}
}

func trimmed(s string) string {
	return strings.TrimSuffix(s, "\n")
}

func main() {
	ports := map[string]string{"Chicago": "8000", "Tokyo": "8010", "UTC": "8020"}
	channels := map[string]chan string{"Chicago": make(chan string), "Tokyo": make(chan string), "UTC": make(chan string)}
	for place, port := range ports {
		go ClockWall(place, port, channels[place])
	}
	fmt.Println("Chicago\tTokyo\tUTC\n")
	ch1 := channels["Chicago"]
	ch2 := channels["Tokyo"]
	ch3 := channels["UTC"]
	for {
		c := <-ch1
		t := <-ch2
		u := <-ch3
		fmt.Printf("loop msg: %s %s %s\n", trimmed(c), trimmed(t), trimmed(u))
	}
	fmt.Scanln()
}
