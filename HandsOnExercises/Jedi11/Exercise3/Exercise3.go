package main

import (
	"fmt"
)

type customErr struct {
	message string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("The error: %v", ce.message)
}

func main() {
	err1 := customErr{
		message: "it no work",
	}
	foo(err1)
}

func foo(e error) {
	fmt.Println(e)
}
