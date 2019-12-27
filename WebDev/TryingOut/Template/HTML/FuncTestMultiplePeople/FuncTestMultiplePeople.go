package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// xUsers is a slice of type User
type xUser struct {
	Users []User
}

// User is the user that tries to log in
type User struct {
	ID         int
	Fname      string
	Lname      string
	Age        int
	Email      string
	Permission bool
}

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.New("tpl.html").Funcs(template.FuncMap{
		"initials": func(user User) string {
			return fmt.Sprintf("%s%s", strings.Split(user.Fname, "")[0], strings.Split(user.Lname, "")[0])
		},
		"dogAge": func(user User) int {
			return user.Age * 7
		},
	}).ParseFiles("tpl.html")

	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		Users := []User{
			{1, "Jeremy", "Nelemans", 17, "jeremynelemans@hotmail.com", true},
			{2, "Tristan", "Goossens", 16, "tristangoossens@gmail.com", true},
			{3, "Henk", "Dubbelman", 65, "hendubbelman@gmail.com", false},
			{4, "Dave", "Bergmans", 28, "db17@rocwb.nl", true},
			{5, "Thomas", "Schrijvernaars", 85, "thomasschrijvernaars@grinder.com", false},
		}
		pass := xUser{Users}
		tpl.ExecuteTemplate(w, "tpl.html", pass)
	})
	http.ListenAndServe(":8080", nil)
}
