package main

import (
	"fmt"
)

var x int

func main() {
	x = 10
	fmt.Println(closure())
}

func closure() int {
	return x
}
