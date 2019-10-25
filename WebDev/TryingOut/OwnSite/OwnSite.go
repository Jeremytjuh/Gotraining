package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type people struct {
	Fname string
	Lname string
	Age   int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	colleagues := []people{
		{Fname: "Jeremy", Lname: "Nelemans", Age: 17},
		{Fname: "Tristan", Lname: "Goossens", Age: 17},
	}
	err := tpl.Execute(os.Stdout, colleagues)
	if err != nil {
		log.Fatalln("An error has occurred")
	}
}
