package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, lierr := net.Listen("tcp", ":8080")
	if lierr != nil {
		log.Fatalln(lierr)
	}
	defer li.Close()

	for {
		con, conerr := li.Accept()
		if lierr != nil {
			log.Println(conerr)
			continue
		}
		go handle(con)
	}
}

func handle(con net.Conn) {
	scanner := bufio.NewScanner(con)
	answer := []string{"h", "i"}
	hangman := make([]string, len(answer))
	for scanner.Scan() {
		game(con, strings.Split(strings.ToLower(scanner.Text()), ""), answer, hangman)
	}
}

func game(con net.Conn, letter, answer, hangman []string) {
	counter := 0
	if len(letter) != 1 {
		fmt.Fprintln(con, "Please only enter 1 letter at a time!")
	} else {
		for i, v := range answer {
			for _, vv := range letter {
				if v == vv {
					hangman[i] = vv
					fmt.Println(hangman)
				} else {
					if len(hangman) == len(answer) {
						for i, v := range hangman {
							if v != answer[i] {
								if counter < 5 {
									if hangman[i] == "" {
										counter++
									} else {
										fmt.Printf("You've won")
									}
								} else {
									fmt.Println("LOST")
								}
							}
						}
					}
				}
			}
		}
		fmt.Println()
	}
}
