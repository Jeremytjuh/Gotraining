package main

import (
	"fmt"
)

func main() {
	start := 0
	end := 10
	fmt.Println("The counter starts at", start, "and it ends at", end, "!")
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
	fmt.Printf("The starting number is of type: %T\n", start)
	fmt.Printf("The ending number is of the type %T/n", end)
}
