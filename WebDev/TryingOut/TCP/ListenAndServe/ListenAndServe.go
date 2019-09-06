package main

import (
	"fmt"
	"net/http"
)

type hotdog int

var d hotdog

// func (m *hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	(*m)++
// 	fmt.Fprintf(w, "%d", *m)
// }

func main() {

	http.HandleFunc("/", funcbla)
	http.HandleFunc("/test", funcbla2)
	http.ListenAndServe(":8080", nil)
}

func funcbla(w http.ResponseWriter, r *http.Request) {
	d++
	fmt.Fprintf(w, "%d", d)
}

func funcbla2(w http.ResponseWriter, r *http.Request) {
	d--
	fmt.Fprintf(w, "%d", d)
}
