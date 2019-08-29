package main

import (
	"fmt"
)

func main() {
	favSport := "Football"
	switch favSport {
	case "American Football":
		fmt.Printf("American Football!!!, my favorite sport is: %v\n", favSport)
	case "Football":
		fmt.Printf("Football!!!, my favorite sport is: %v\n", favSport)
	case "Baseball":
		fmt.Printf("Baseball!!!, my favorite sport is: %v\n", favSport)
	case "Tennis":
		fmt.Printf("Tennis!!!, my favorite sport is: %v\n", favSport)
	}
}
