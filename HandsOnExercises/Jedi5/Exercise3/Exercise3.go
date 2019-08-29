package main

import (
	"fmt"
)

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	a := truck{
		vehicle: vehicle{
			door:  6,
			color: "blue",
		},
		fourWheel: true,
	}

	b := sedan{
		vehicle: vehicle{
			door:  4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("The amount of doors on the truck is %v\n", a.door)
	fmt.Printf("The amount of doors on the sedan is %v\n", b.door)
}
