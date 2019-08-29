package main

import (
	"fmt"
)

var a int

func main() {
	a = 69
	fmt.Printf("%d\t%b\t\t%#x\t%T\n", a, a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\t%T\n", b, b, b, b)
}
