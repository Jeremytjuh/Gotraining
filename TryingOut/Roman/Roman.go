package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	filename := args()
	checkFile(filename)
	content := scanInput()
	writeToFile(filename, content)
}

// args checks whether you have entered the correct amount of arguments
func args() string {
	xargs := os.Args
	if len(xargs) != 2 {
		log.Fatalf("Invalid amount of arguments. Expected 2, Got %d", len(xargs))
	}
	return xargs[1]
}

// checkFile checks whether the given file exists
func checkFile(filename string) {
	if !strings.HasSuffix(filename, ".md") {
		filename = fmt.Sprintf("%s.md", filename)
	}

	_, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("File %s does not exist", filename)
	}
}

// scanInput scans the os.Stdin input
func scanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	return scanner.Text()
}

func writeToFile(filename, content string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open file %s", filename)
	}

}
