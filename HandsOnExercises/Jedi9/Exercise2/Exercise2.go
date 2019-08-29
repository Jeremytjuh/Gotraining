package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello, I am a person:", p)
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		name: "Jeremy Nelemans",
	}

	p2 := person{
		name: "Tristan Goossens",
	}
	saySomething(&p1) //This works
	saySomething(&p2) //This works
	// saySomething(p1) This doesn't work
}

func saySomething(h human) {
	h.speak()
}
