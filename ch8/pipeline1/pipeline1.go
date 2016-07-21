// Pipeline program that sends data through channels
package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 1st stage - counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// 2nd stage - squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// 3rd stage - Printer
	for {
		fmt.Println(<-squares)
		time.Sleep(100 * time.Millisecond)
	}
}
