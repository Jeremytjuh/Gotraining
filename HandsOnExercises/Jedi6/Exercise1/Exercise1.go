package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 69
}

func bar() (int, string) {
	return 69, "420"
}
