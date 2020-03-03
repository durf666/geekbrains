package main

import (
	"fmt"
	// "sync"
	// "time"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 1000; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
		// time.Sleep (100 * time.Millisecond)

	}

}
