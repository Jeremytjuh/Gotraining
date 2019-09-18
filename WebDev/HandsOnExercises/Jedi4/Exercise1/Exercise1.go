package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("tpl.html")
	if err != nil {
		log.Println("Parse error: ", err)
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Tfoe ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
	tpl.ExecuteTemplate(w, "tpl.html", nil)
}
