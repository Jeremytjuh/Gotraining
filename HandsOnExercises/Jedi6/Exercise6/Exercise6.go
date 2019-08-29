package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("This was printed using an anonymous function")
	}()
}
