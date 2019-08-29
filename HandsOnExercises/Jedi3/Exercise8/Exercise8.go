package main

import (
	"fmt"
)

func main() {
	truee := true
	falsee := false
	switch {
	case truee:
		{
			fmt.Printf("The value is True\n")
		}
	case falsee:
		{
			fmt.Printf("The value is False\n")
		}
	default:
		fmt.Printf("The value is not True nor False\n")
	}
}
