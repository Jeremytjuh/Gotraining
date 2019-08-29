package main

import (
	"fmt"
)

func main() {
	birthyear := 2002 //Depends on what your opinion is on when you're alive
	yearnow := 2019
	for i := birthyear; i <= yearnow; i++ {
		fmt.Println(i)
	}
}
