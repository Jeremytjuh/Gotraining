package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	favIcecream []string
}

func main() {
	Jeremy := person{
		first:       "Jeremy",
		last:        "Nelemans",
		favIcecream: []string{"Oreo", "Chocolate"},
	}

	Tristan := person{
		first:       "Tristan",
		last:        "Goossens",
		favIcecream: []string{"Strawberry", "Vanilla"},
	}

	fmt.Println(Jeremy.first, Jeremy.last)
	for i, v := range Jeremy.favIcecream {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("\n")
	fmt.Println(Tristan.first, Tristan.last)
	for i, v := range Tristan.favIcecream {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("\n")
}
