package main

import (
	"html/template"
	"log"
	"net/http"
)

// ViewData is the data that you can view
type ViewData struct {
	User User
}

// User is the user that tries to log in
type User struct {
	ID    int
	Email string
}

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.New("tpl.html").Funcs(template.FuncMap{
		"hasPermission": func(user User, feature string) bool {
			if user.ID == 1 && feature == "feature-a" {
				return true
			}
			return false
		},
	}).ParseFiles("tpl.html")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conten-Type", "text/html")
		user := User{
			ID:    1,
			Email: "jon@ahneef.com",
		}
		vd := ViewData{user}
		tpl.Execute(w, vd)
	})
	http.ListenAndServe(":8080", nil)
}
