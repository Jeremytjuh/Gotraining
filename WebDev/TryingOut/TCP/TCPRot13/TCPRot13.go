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
		log.Fatalln("Fatal connection error occurred: ", lierr)
	}
	defer li.Close()

	for {
		con, conerr := li.Accept()
		if conerr != nil {
			log.Println(conerr)
			continue
		}
		go handle(con)
	}
}

func handle(con net.Conn) {
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		li := strings.ToLower(scanner.Text())
		bs := []byte(li)
		r := rot13(bs)

		fmt.Fprintf(con, "%s - %s\n\n", li, r)
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
