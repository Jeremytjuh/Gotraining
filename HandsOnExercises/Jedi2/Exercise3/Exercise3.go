package main

import (
	"fmt"
)

const (
	a        = 420
	b int    = 69
	c string = "Hello"
)

func main() {
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
}
