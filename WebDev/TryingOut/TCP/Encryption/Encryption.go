package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	// "strings"
)

var counter int
var algo int

func init() {
	counter = 1
}

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
		raw := []byte(scanner.Text())
		encrypted := encryptor(raw)
		decrypted := decryptor(encrypted)
		fmt.Fprintf(con, "Encrypted: %s\nDecrypted: %s\n\n", encrypted, decrypted)
		fmt.Printf("[%d]Algo:\t%d\tEncrypted:%s\n", counter, algo, encrypted)
		counter++
	}
	defer con.Close()
}

func encryptor(xb []byte) []byte {
	length := len(xb)
	var xbb = make([]byte, length)
	rand.Seed(int64(length))
	algo = rand.Intn(20)
	for {
		if algo == 0 {
			algo = rand.Intn(length)
		} else {
			break
		}
	}
	if length%2 == 0 {
		for i, v := range xb {
			if i%2 == 0 {
				xbb[i] = v + byte(algo)
			} else {
				xbb[i] = v - byte(algo)
			}
		}
	} else {
		for i, v := range xb {
			if i%2 == 0 {
				xbb[i] = v - byte(algo)
			} else {
				xbb[i] = v + byte(algo)
			}
		}
	}
	return xbb
}

func decryptor(xb []byte) []byte {
	length := len(xb)
	var xbb = make([]byte, length)
	rand.Seed(int64(length))
	algo := rand.Intn(20)
	for {
		if algo == 0 {
			algo = rand.Intn(length)
		} else {
			break
		}
	}
	if length%2 == 0 {
		for i, v := range xb {
			if i%2 == 0 {
				xbb[i] = v - byte(algo)
			} else {
				xbb[i] = v + byte(algo)
			}
		}
	} else {
		for i, v := range xb {
			if i%2 == 0 {
				xbb[i] = v + byte(algo)
			} else {
				xbb[i] = v - byte(algo)
			}
		}
	}
	return xbb
}
