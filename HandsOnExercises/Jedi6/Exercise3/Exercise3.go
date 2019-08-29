package main

import (
	"fmt"
)

func main() {
	defer foo()
}

func foo() {
	defer func() {
		fmt.Println("Boven")
	}()
	fmt.Println("Onder")
}
