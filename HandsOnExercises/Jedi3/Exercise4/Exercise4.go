package main

import (
	"fmt"
)

func main() {
	birthyear := 2002
	yearnow := 2019
	for {
		if birthyear > yearnow {
			break
		}
		fmt.Printf("%v\n", birthyear)
		birthyear++
	}
}
