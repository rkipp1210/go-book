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
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// 2nd stage - squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// 3rd stage - Printer
	for x := range squares {
		fmt.Println(x)
		time.Sleep(10 * time.Millisecond)
	}
}
