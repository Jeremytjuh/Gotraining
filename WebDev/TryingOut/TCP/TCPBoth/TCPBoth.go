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
		log.Fatalln("Listen error occurred: ", listenerr)
	}
	defer listen.Close()

	for {
		connection, conerr := listen.Accept()
		if conerr != nil {
			log.Println("Connection error occurred: ", conerr)
			continue
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Client said: %s\n", line)
		switch line {
		case "Hello":
			fmt.Fprint(connection, "Server said: Hello! How are you doing?\n")
		case "Bad":
			fmt.Fprint(connection, "That's not good. How come?\n")
		case "Bad...":
			fmt.Fprint(connection, "That's not good. How come?\n")
		case "Great":
			fmt.Fprint(connection, "Amazing!\n")
		case "Great!":
			fmt.Fprint(connection, "Amazing!\n")
		case "Amazing":
			fmt.Fprint(connection, "Wonderful!\n")
		case "Amazing!":
			fmt.Fprint(connection, "Wonderful!\n")
		case "How are you doing?":
			fmt.Fprint(connection, "Server said: I'm doing great! Thanks for asking!\n")
		default:
			fmt.Fprint(connection, "Server said: I did not catch that, could you please rephrase that?\n")
		}
		// fmt.Fprintf(connection, "I heard you say: %s\n", line)
	}
	defer connection.Close()
}
