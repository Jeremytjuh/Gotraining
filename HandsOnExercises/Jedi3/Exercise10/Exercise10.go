package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The resulting bool value of %v and %v = %v\n", true, true, (true && true))
	fmt.Printf("The resulting bool value of %v and %v = %v\n", true, false, (true && false))
	fmt.Printf("The resulting bool value of %v or %v = %v\n", true, true, (true || true))
	fmt.Printf("The resulting bool value of %v or %v = %v\n", true, false, (true || false))
	fmt.Printf("The resulting bool value of negative %v = %v\n", true, (!true))
}
