package main

import (
	"fmt"
	"html/template"
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
	tpl = template.Must(template.New("tpl.html").Funcs(template.FuncMap{
		"hasPermission": func(user User) bool {
			if user.Permission == true {
				return true
			}
			return false
		},
		"initials": func(user User) string {
			xFirst := strings.Split(user.Fname, "")
			xLast := strings.Split(user.Lname, "")
			return fmt.Sprintf("%s%s", xFirst[0], xLast[0])
		},
		"dogAge": func(user User) int {
			return user.Age * 7
		},
	}).ParseFiles("tpl.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conten-Type", "text/html")
		Users := []User{
			User{1, "Jeremy", "Nelemans", 17, "jeremynelemans@hotmail.com", true},
			User{2, "Tristan", "Goossens", 17, "tristangoossens@gmail.com", true},
			User{3, "Henk", "Dubbelman", 65, "hendubbelman@gmail.com", false},
			User{4, "Dave", "Bergmans", 28, "db17@rocwb.nl", true},
		}
		pass := xUser{Users}
		tpl.ExecuteTemplate(w, "tpl.html", pass)
	})
	http.ListenAndServe(":8080", nil)
}
