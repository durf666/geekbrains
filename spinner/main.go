package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(150 * time.Millisecond)
	const n = 42
	time.Sleep(10 * time.Second)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
