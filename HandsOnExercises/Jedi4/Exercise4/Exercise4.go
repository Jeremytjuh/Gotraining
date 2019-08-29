package main

import (
	"fmt"
)

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	b := []int{52}
	c := append(a, b...)
	fmt.Println(c)
	d := []int{53, 54, 55}
	e := append(c, d...)
	fmt.Println(e)
	f := []int{56, 57, 58, 59, 60}
	g := append(e, f...)
	fmt.Println(g)
}
