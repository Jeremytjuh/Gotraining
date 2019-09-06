package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var counter int

func main() {
	counter = 1
	li, lierr := net.Listen("tcp", ":8080")
	if lierr != nil {
		log.Fatalln("Fatal connection error: ", lierr)
	}
	defer li.Close()

	for {
		con, conerr := li.Accept()
		if conerr != nil {
			log.Println("Conerr: ", conerr)
			continue
		}
		go handle(con)
	}
}

func handle(con net.Conn) {
	li := bufio.NewScanner(con)
	for li.Scan() {
		line := strings.ToLower(li.Text())
		ss := strings.Split(line, "")
		if len(ss) != 0 {
			fmt.Printf("Input %d: %s\n", counter, line)
			counter++
			for i, v := range ss {
				switch v {
				case "a":
					fmt.Fprintf(con, ".- ")
				case "b":
					fmt.Fprintf(con, "-... ")
				case "c":
					fmt.Fprintf(con, "-.-. ")
				case "d":
					fmt.Fprintf(con, "-.. ")
				case "e":
					fmt.Fprintf(con, ". ")
				case "f":
					fmt.Fprintf(con, "..-. ")
				case "g":
					fmt.Fprintf(con, "--. ")
				case "h":
					fmt.Fprintf(con, ".... ")
				case "i":
					fmt.Fprintf(con, ".. ")
				case "j":
					fmt.Fprintf(con, ".--- ")
				case "k":
					fmt.Fprintf(con, "-.- ")
				case "l":
					fmt.Fprintf(con, ".-.. ")
				case "m":
					fmt.Fprintf(con, "-- ")
				case "n":
					fmt.Fprintf(con, "-. ")
				case "o":
					fmt.Fprintf(con, "--- ")
				case "p":
					fmt.Fprintf(con, ".--. ")
				case "q":
					fmt.Fprintf(con, "--.- ")
				case "r":
					fmt.Fprintf(con, ".-. ")
				case "s":
					fmt.Fprintf(con, "... ")
				case "t":
					fmt.Fprintf(con, "- ")
				case "u":
					fmt.Fprintf(con, "..- ")
				case "v":
					fmt.Fprintf(con, "...- ")
				case "w":
					fmt.Fprintf(con, ".-- ")
				case "x":
					fmt.Fprintf(con, "-..- ")
				case "y":
					fmt.Fprintf(con, "-.-- ")
				case "z":
					fmt.Fprintf(con, "--.. ")
				case " ":
					fmt.Fprintf(con, "/ ")
				case "0":
					fmt.Fprintf(con, "----- ")
				case "1":
					fmt.Fprintf(con, ".---- ")
				case "2":
					fmt.Fprintf(con, "..--- ")
				case "3":
					fmt.Fprintf(con, "...-- ")
				case "4":
					fmt.Fprintf(con, "....- ")
				case "5":
					fmt.Fprintf(con, "..... ")
				case "6":
					fmt.Fprintf(con, "-.... ")
				case "7":
					fmt.Fprintf(con, "--... ")
				case "8":
					fmt.Fprintf(con, "---.. ")
				case "9":
					fmt.Fprintf(con, "----. ")
				default:
					fmt.Fprintf(con, "? ")
				}
				if i == (len(ss) - 1) {
					fmt.Fprintf(con, "\n")
				}
			}
		}
	}
}
