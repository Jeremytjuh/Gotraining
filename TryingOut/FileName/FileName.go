package main

import (
	"filepath"
	"log"
	"runtime"
)

func main() {
	fmt.Println(getFileName())
}

func getFileName() string {
	_, fpath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Failed to get filename")
	}
	filename := filepath.Base(fpath)
	// remove extension
	filename = removeExtension(filename)
	return filename + ".log"
}
