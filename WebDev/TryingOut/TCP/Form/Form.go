package main

import (
	"fmt"
	// "html/template"
	"net/http"
)

type counter int

// var tpl *template.Template

var i counter

// func init() {
// 	tpl = template.Must(template.ParseFiles("tpl.html"))
// }

func main() {
	http.HandleFunc("/increase", increase)
	http.HandleFunc("/decrease", decrease)
	http.ListenAndServe(":8080", nil)
}

func increase(w http.ResponseWriter, req *http.Request) {
	i++
	fmt.Fprintf(w, "%d", i)
}

func decrease(w http.ResponseWriter, req *http.Request) {
	i--
	fmt.Fprintf(w, "%d", i)
}

// func (i counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	err := req.ParseForm()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	tpl.ExecuteTemplate(w, "tpl.html", req.Form)
// }
