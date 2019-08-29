package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "Jer"
	p.last = "Nel"
	p.age = 42
}

func main() {
	Jeremy := person{
		first: "Jeremy",
		last:  "Nelemans",
		age:   17,
	}
	fmt.Println(Jeremy)
	changeMe(&Jeremy)
	fmt.Println(Jeremy)
}
