package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	con, conerr := net.Dial("tcp", "localhost:8080")
	if conerr != nil {
		log.Fatalln("Connection error occurred: ", conerr)
	}
	defer con.Close()

	bs, bserr := ioutil.ReadAll(con)
	if bserr != nil {
		log.Println("Read error occurred: ", bserr)
	}

	fmt.Println(string(bs))
}
