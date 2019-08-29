package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
	xown := []string{"c", "f", "e", "a", "d", "b"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Printf("%v\n\n", xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Printf("%v\n\n", xs)

	fmt.Println(xown)
	sort.Strings(xown)
	fmt.Printf("%v\n\n", xown)
}
