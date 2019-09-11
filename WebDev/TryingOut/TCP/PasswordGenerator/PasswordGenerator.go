package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	li, lierr := net.Listen("tcp", ":8080")
	if lierr != nil {
		log.Fatalln(lierr)
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
		length, lengtherr := strconv.ParseInt(scanner.Text(), 10, 64)
		if lengtherr != nil {
			fmt.Fprintln(con, "Error: Not a valid number")
			continue
		}
		if length < 2147483647 {
			fmt.Printf("Length: %d\n", length)
			password := genpass(length)
			fmt.Printf("Password: %s\n", password)
			fmt.Fprintf(con, "Password: %s\n", password)
		} else {
			fmt.Fprintln(con, "Error: Number too high")
		}
	}
}

func genpass(length int64) string {
	xi := make([]rune, length)
	for i := range xi {
		rand.Seed(time.Now().UnixNano())
		rdnb := rand.Intn(126)
		for rdnb < 33 || rdnb > 126 {
			rand.Seed(time.Now().UnixNano())
			rdnb = rand.Intn(126)
		}
		xi[i] = rune(rdnb)
	}
	var pass string
	for _, v := range xi {
		pass = pass + string(v)
	}
	return pass
}
