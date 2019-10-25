package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, lierr := net.Listen("tcp", ":8080")
	if lierr != nil {
		log.Fatalln("Fatal listen error occurred: ", lierr)
	}

	defer li.Close()

	for {
		con, conerr := li.Accept()
		if conerr != nil {
			log.Println("Fatal connection error occurred: ", conerr)
			continue
		}
		go handle(con)
	}
}

func handle(con net.Conn) {
	err := con.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("Fatal connection timeout", err)
	}

	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(con, "I heard you say: %s", line)
	}

	defer con.Close()

	fmt.Println("Code got here*********************************")
}
