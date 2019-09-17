package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var docscanner *bufio.Scanner
var pcscanner *bufio.Scanner

func init() {
	var err error
	tpl, err = template.New("LoginTest.proto").Funcs(template.FuncMap{
		"import": importMarkdown,
	}).ParseFiles("LoginTest/LoginTest/LoginTest.proto")

	if err != nil {
		log.Println("Parse error: ", err)
	}
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Println("Execute error: ", err)
	}
}

func importMarkdown(name string) string {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		log.Println(err)
	}

	return string(file)
}
