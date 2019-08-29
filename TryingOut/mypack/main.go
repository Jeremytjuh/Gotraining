package main

import (
	"fmt"
	"github.com/Jeremytjuh/Gotraining/TryingOut/mypack/mypack"
	"math"
)

func main() {
	xi := []int{2, 7, 1, 3, 66, 32, 23}
	xii := mypack.SliceIntToFloat64(xi)
	for _, v := range xii {
		fmt.Printf("The area of a circle with a radius of %v cm is %v cm\n", v, v*v*math.Pi)
	}
}
