package ch8_goroutines

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// deliberately inefficient implementation
func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}

/*
The main goroutine computes the 45th Fibonacci number. Since this is slow, a
spinner provides the user with a visual indication that the program is still
running. The spinner runs in its own goroutine. Once the fibonacci computation
returns, all goroutines terminate because the main function has finished.
*/
