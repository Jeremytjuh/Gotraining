package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("File already exists, do you want to overwrite it? WARNING: THIS ACTION IS IRREVERSIBLE\nYes or No? ")
	scanner := bufio.NewScanner(os.Stdin)
Loop:
	for scanner.Scan() {
		switch strings.ToLower(scanner.Text()) {
		case "yes":
			fmt.Println("You chose yes!")
			executeFurther()
			break Loop
		case "no":
			log.Fatalln("Quit program: chose no")
		default:
			log.Fatalln("Quit program: invalid command")
		}
	}
}

func executeFurther() {
	fmt.Println("Further executing!")
}
