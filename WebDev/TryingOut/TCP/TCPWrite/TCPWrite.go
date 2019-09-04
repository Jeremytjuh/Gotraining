package main

import (
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
			log.Fatalln("Connection error occured: ", conerr)
			continue
		}

		fmt.Fprintln(connection, "Ewa")
		connection.Close()
	}
}
