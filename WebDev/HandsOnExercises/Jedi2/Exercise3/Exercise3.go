package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("tpl.html")
	if err != nil {
		log.Println("Parse error : ", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(root))
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "tpl.html", "Oi cunt")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "tpl.html", "Oi Dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "tpl.html", "Oi Jeremy")
}
