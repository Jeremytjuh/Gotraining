package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Root")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Oi cunt")
}
