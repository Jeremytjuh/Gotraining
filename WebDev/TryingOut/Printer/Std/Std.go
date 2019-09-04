package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	genNums(1, 100)
	sepNums("Even.txt", "Odd.txt")
}

func genNums(start, end int) {
	Counters, _ := os.Create("Counter.txt")
	os.Stdout = Counters
	for i := start; i < end; i++ {
		fmt.Println(i)
	}
}

func sepNums(even, odd string) {
	ogstd := bufio.NewScanner(os.Stdin)
	i := 0
	Even, _ := os.Create(even)
	Odd, _ := os.Create(odd)
	for ogstd.Scan() {
		if i%2 != 0 {
			os.Stdout = Even
			fmt.Printf("%s\n", ogstd.Text())
		} else {
			os.Stdout = Odd
			fmt.Printf("%s\n", ogstd.Text())
		}
		i++
	}
}
