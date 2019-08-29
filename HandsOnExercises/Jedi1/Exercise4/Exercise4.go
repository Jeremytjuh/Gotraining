package main

import (
	"fmt"
)

type jer int

var x jer

func main() {
	fmt.Printf("Value: %v\n", x)
	fmt.Printf("Type: %T\n", x)
	x = 42
	fmt.Printf("Value: %v\n", x)
}
