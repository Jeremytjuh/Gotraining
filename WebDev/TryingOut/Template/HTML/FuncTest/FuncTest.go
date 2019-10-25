package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

// ViewData is a struct to view data
type ViewData struct {
	User User
}

// User is a struct of user information
type User struct {
	ID    int
	Email string
}

func main() {
	var err error
	testTemplate, err = template.New("tpl.html").Funcs(template.FuncMap{
		"hasPermission": func(user User, feature string) bool {
			if user.ID == 1 && feature == "feature-a" {
				return true
			}
			return false
		},
	}).ParseFiles("tpl.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	user := User{
		ID:    1,
		Email: "jon@calhoun.io",
	}
	vd := ViewData{user}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
