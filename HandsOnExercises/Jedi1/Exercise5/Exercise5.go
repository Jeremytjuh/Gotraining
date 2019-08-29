package main

import (
	"fmt"
)

type jer int

var x jer

var y int

func main() {
	fmt.Printf("x Value: %v\n", x)
	fmt.Printf("x Type: %T\n", x)
	x = 42
	fmt.Printf("x Value: %v\n\n", x)

	y = int(x)

	fmt.Printf("y Value: %v\n", y)
	fmt.Printf("y Type: %T\n", y)
}
