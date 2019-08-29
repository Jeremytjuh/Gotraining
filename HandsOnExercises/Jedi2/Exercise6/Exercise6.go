package main

import (
	"fmt"
)

const (
	year1 = 2017 + iota
	year2 = 2017 + iota
	year3 = 2017 + iota
	year4 = 2017 + iota
)

func main() {
	fmt.Printf("It is now year number: %v, this number is of the type: %T\n", year1, year1)
	fmt.Printf("It is now year number: %v, this number is of the type: %T\n", year2, year2)
	fmt.Printf("It is now year number: %v, this number is of the type: %T\n", year3, year3)
	fmt.Printf("It is now year number: %v, this number is of the type: %T\n", year4, year4)
}
