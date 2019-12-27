package main

import (
	"io/ioutil"
	"log"
	"time"
)

func main() {
	dir, err := ioutil.TempDir(".", "testdir")
	if err != nil {
		log.Fatalln(err)
	}

}
