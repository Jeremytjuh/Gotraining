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
		favIcecream: []string{"Strawberry", "Vanilla", "Cookies"},
	}

	a := map[string]person{
		Jeremy.last:  Jeremy,
		Tristan.last: Tristan,
	}

	for k, v := range a {
		fmt.Printf("%v %v's key is: %v\nHis favorite icecream flavours are:\n", v.first, v.last, k)
		for i, v2 := range v.favIcecream {
			fmt.Printf("%v\t%v\n", i+1, v2)
		}
		fmt.Printf("\n")
	}
}
