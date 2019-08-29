package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func (s square) area() float64 {

	return (s.length * s.width)
}

func (c circle) area() float64 {
	return (math.Sqrt(c.radius) * math.Pi)
}

func main() {
	square1 := square{
		length: 10,
		width:  10,
	}

	circle1 := circle{
		radius: 5,
	}

	info(square1)
	info(circle1)
}
