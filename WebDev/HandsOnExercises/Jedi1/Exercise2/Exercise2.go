package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	pass := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "Southern",
		},

		hotel{
			Name:    "Bastion",
			Address: "69 Street",
			City:    "Roosendaal",
			Zip:     "696969 AB",
			Region:  "Central",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", pass)
	if err != nil {
		log.Fatalln(err)
	}
}
