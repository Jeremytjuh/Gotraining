package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("StartingFiles/Public/Pics/tpl.html"))
}

func main() {
	fs := http.FileServer(http.Dir("Public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", execute)
	http.ListenAndServe(":8080", nil)
}

func execute(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("Fatal Execute error occurred!, %s", err)
	}
}
