package main

import (
	"io/ioutil"
	"log"
	"net"
	"fmt"
)

func main() {
	con, conerr := net.Dial("tcp", "localhost:8080")
	if conerr != nil {
		log.Fatalln("Connection error occured: ", conerr)
	}
	defer con.Close()

	bs, bserr := ioutil.ReadAll(con)
	if bserr != nil {
		log.Println("Read error occured: ", bserr)
	}

	fmt.Println(string(bs))
}
