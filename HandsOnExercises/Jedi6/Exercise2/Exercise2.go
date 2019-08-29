package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(foo(a...))
	b := []int{}
	fmt.Println(bar(b...))
}

func foo(a ...int) int {
	b := 0
	for _, v := range a {
		b += v
	}
	return b
}

func bar(a ...int) int {
	b := 0
	for _, v := range a {
		b += v
	}
	return b
}
