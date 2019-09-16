package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tpl.Execute(w, nil)
			return
		}
		firstname := r.FormValue("fname")
		lastname := r.FormValue("lname")

		tpl.ExecuteTemplate(w, "tpl.html", struct {
			Firstname string
			Lastname  string
		}{firstname, lastname})
	})
	http.ListenAndServe(":8080", nil)
}
