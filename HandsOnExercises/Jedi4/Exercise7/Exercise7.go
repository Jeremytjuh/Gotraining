package main

import (
	"fmt"
)

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenney", "Helloooooo, James"}
	c := [][]string{a, b}
	fmt.Println("Multidimensional slice:")
	for i, v := range c {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("\nJames Bond slice:\n")
	for i, v := range a {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("\nMiss Moneypenney slice:\n")
	for i, v := range b {
		fmt.Printf("%v\t%v\n", i, v)
	}
}
