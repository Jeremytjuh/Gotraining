package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, listenerr := net.Listen("tcp", ":8080")
	if listenerr != nil {
		log.Fatalln("Listen error occured: ", listenerr)
	}
	defer listen.Close()

	for {
		connection, conerr := listen.Accept()
		if conerr != nil {
			log.Println("Connection error occured: ", conerr)
			continue
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer connection.Close()
}
