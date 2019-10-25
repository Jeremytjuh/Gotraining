package main

import (
	"fmt"
)

func main() {
	g := (1 == 11)
	h := (1 <= 1)
	i := (1 >= 1)
	j := (1 != 1)
	k := (1 < 1)
	l := (1 > 1)
	fmt.Printf("%v\t%T\n", g, g)
	fmt.Printf("%v\t%T\n", h, h)
	fmt.Printf("%v\t%T\n", i, i)
	fmt.Printf("%v\t%T\n", j, j)
	fmt.Printf("%v\t%T\n", k, k)
	fmt.Printf("%v\t%T\n", l, l)
}
