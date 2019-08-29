package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for a := 1; a <= 3; a++ {
			fmt.Printf("%#U\n", i)
		}
		fmt.Printf("\n")
	}
}
