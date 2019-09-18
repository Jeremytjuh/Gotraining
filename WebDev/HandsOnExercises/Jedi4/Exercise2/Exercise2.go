package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatalf("Fatal error occured: %s", http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
