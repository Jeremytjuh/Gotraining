package main

import (
	"fmt"
)

func main() {
	var answer1 string
	fmt.Println("Name: ")
	a, _ := fmt.Scan(&answer1)
	switch {
	case a == 1:
		fmt.Println("Yes")
	default:
		fmt.Println("No")
	}
}
