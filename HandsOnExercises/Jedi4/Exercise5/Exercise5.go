package main

import (
	"fmt"
)

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	b := a[:3]
	c := a[6:]
	d := append(b, c...)
	fmt.Println(d)
}
