package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, lierr := net.Listen("tcp", ":8080")
	if lierr != nil {
		log.Println(lierr)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		io.WriteString(con, "I see you connected eh?\n")
		con.Close()
	}
}
