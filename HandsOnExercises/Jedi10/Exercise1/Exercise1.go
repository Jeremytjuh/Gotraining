package main

import (
	"fmt"
)

func main() {
	selfexec()
	buffchan()
}

func selfexec() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func buffchan() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
