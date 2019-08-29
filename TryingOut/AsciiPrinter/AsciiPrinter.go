package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Dec\tHex\tUni\n")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\n", i, i, i)
	}
}
