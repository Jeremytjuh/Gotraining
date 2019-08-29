package main

import (
	"fmt"
)

func main() {
	a := struct {
		first string
		last  string
	}{
		first: "Jer",
		last:  "Nel",
	}
	fmt.Println(a)
}
