package main

import (
	"fmt"
	"github.com/Jeremytjuh/golang/Dog"
	"github.com/Jeremytjuh/golang/Swapper"
)

func main() {
	first := "Jeremy"
	last := "Nelemans"
	swapped := fmt.Sprintln(swapper.Swapper(first, last))
	fmt.Printf("%v %v swapped around is %v\n", first, last, swapped)
	fmt.Printf("17 in dog years is: %v\n", dog.Humantodog(17))
}
