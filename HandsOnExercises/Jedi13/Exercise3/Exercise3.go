package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/03/01-code-starting/mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := [][]int{a, b, c, d, e}
	return f
}
