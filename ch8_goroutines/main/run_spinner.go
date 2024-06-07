package main

import (
	"fmt"
	ch8 "github.com/megulus/go-book/ch8_goroutines"
	"time"
)

/*
The main goroutine computes the 45th Fibonacci number. Since this is slow, a
spinner provides the user with a visual indication that the program is still
running. The spinner runs in its own goroutine. Once the fibonacci computation
returns, all goroutines terminate because the main function has finished.
*/

func main() {
	go ch8.Spinner(100 * time.Millisecond)
	const n = 45
	fibN := ch8.Fib(n) // this is slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
