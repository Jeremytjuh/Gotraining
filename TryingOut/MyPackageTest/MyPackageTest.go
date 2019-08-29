package main

import (
	"fmt"
)

func main() {
	fmt.Println(pSwapper("Tristan", "Goossens"))
}

func pSwapper(s1 string, s2 string) (string, string) {
	s1, s2 = s2, s1
	return s1, s2
}
