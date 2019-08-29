package main

import (
	"fmt"
)

func main() {
	a := funcception()
	fmt.Println(a())
}

func funcception() func() int {
	return func() int {
		return 21
	}
}
