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
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
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
