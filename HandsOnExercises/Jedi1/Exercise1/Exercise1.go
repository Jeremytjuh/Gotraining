package main

import (
	"fmt"
)

func main() {
	x := 42
	y := `"James Bond"`
	z := true
	fmt.Println(x, y, z)
	fmt.Printf("x is van het type %T en zijn waarde is %v\n", x, x)
	fmt.Printf("y is van het type %T en zijn waarde is %v\n", y, y)
	fmt.Printf("z is van het type %T en zijn waarde is %v\n", z, z)
}
