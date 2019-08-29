package main

import (
	"fmt"
)

func main() {
	a := []int{2, 7, 43, 565, 22, 342, 65656, 22, 1, 5435}
	for i, v := range a {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("%T\n", a)
}
