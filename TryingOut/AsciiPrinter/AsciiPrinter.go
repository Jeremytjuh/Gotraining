package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		convert(i)
	}
}

func loop() {
	fmt.Printf("Dec\tHex\tUni\n")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\n", i, i, i)
	}
}

func convert(convert int) {
	fmt.Printf("%d\t%#U\n", convert, convert)
}
