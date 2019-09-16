package main

import (
	"bufio"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var docscanner *bufio.Scanner
var pcscanner *bufio.Scanner

func init() {
	tpl = template.Must(template.New("docgen").Funcs(template.FuncMap{
		"import": func(name string) string {
			var retval string
			for docscanner.Scan() {
				retval += (docscanner.Text())
			}
			return retval
		},
	}).ParseFiles("LoginTest/LoginTest/LoginTest.proto"))
}

func main() {
	docscanner, pcscanner = scanners("Documentation/documentation.md", "LoginTest/LoginTest/LoginTest.proto")
}

func scanners(documentation, protofile string) (*bufio.Scanner, *bufio.Scanner) {
	docreader, docreadererr := os.Open(documentation)
	if docreadererr != nil {
		log.Println("Doc read error occured: ", docreadererr)
	}

	pcreader, pcreadererr := os.Open(protofile)
	if pcreadererr != nil {
		log.Println("Proto read error occured: ", pcreadererr)
	}
	return bufio.NewScanner(docreader), bufio.NewScanner(pcreader)
}
