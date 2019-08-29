package main

import (
	"fmt"
)

func main() {
	a := [5]int{5, 9, 21, 12, 85}
	for i, v := range a {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("Type:\t%T\n", a)
}
