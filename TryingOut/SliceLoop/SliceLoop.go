package main

import (
	"fmt"
)

func main() {
	x := []int{8, 3, 1, 7, 20}
	for i, v := range x {
		fmt.Printf("On position %v the value %v is stored\n", i, v)
	}
}
