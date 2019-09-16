package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	docscanner, pcscanner := scanners("documentation.md", "LoginTest/LoginTest/LoginTest.proto")
	var insert bool
	for docscanner.Scan() {
		docline := docscanner.Text()
		if insert == true {
			for docline != "{{end}}" {
				fmt.Printf("// %s\n", docline)
				if docscanner.Scan() {
					docline = docscanner.Text()
				} else {
					log.Fatalln("ERROR: Missing end tag, reached end of file")
				}
			}
			insert = false
			continue
		}
		for pcscanner.Scan() {
			pcline := pcscanner.Text()
			if fmt.Sprintf("// %s", docline) == pcline {
				insert = true
				break
			} else {
				fmt.Println(pcline)
			}
		}
	}
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
