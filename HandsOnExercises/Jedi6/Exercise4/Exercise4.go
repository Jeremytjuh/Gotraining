package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello, my name is %v %v, i'm %v years old", p.first, p.last, p.age)
}

func main() {
	Jeremy := person{
		first: "Jeremy",
		last:  "Nelemans",
		age:   17,
	}

	Tristan := person{
		first: "Tristan",
		last:  "Goossens",
		age:   17,
	}
	Jeremy.speak()
	fmt.Printf("\n")
	Tristan.speak()
	fmt.Printf("\n")
}
