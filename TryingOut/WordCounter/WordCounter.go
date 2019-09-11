package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for k, v := range wordcounter(scanner.Text()) {
			fmt.Printf("%s\t%d\n", k, v)
		}
	}
}

func wordcounter(s string) map[string]int {
	charlist := strings.Split(s, "")
	amount := make(map[string]int)
	for _, word := range charlist {
		_, ok := amount[word]
		if ok {
			amount[word]++
		} else {
			amount[word] = 1
		}
	}
	return amount
}
